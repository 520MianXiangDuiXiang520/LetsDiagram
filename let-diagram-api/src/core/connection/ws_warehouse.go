package connection

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"lets_diagram/src"
	"lets_diagram/src/core/connection/pingpong"
	"log"
	"sync"
	"time"
)

// Wheat 表示一个参与协作的 WebSocket 连接
type Wheat struct {
	Conn          *websocket.Conn // webSocket 连接对象
	UserID        uint            // 该连接对应的用户
	LatestVersion int64           // 最后一次同步给该连接的版本号
}

func (w *Wheat) Notify(data []byte) error {
	return w.Conn.WriteMessage(websocket.TextMessage, data)
}

type Granary struct {
	WheatList []*Wheat // 该 Canvas 对应的所有连接
	CanvasID  uint     // 该连接对应的 Canvas
	PCU       int
	// 当画布的创作者断开连接后, 所有协作者等待一段时间，如果创作者没有重新加入，终止此次协作
	// 等待期间，画布被锁定，所有人无法操作。
	RootCloseWaiter *time.Timer
	// 画布创作者断开后重连，触发该信号，该信号会重置 rootCloseWaiter 并继续允许正常协作
	RootRejoinSign chan struct{}
}

func (g *Granary) Find(user uint) (*Wheat, bool) {
	for _, wheat := range g.WheatList {
		if wheat.UserID == user {
			return wheat, true
		}
	}
	return nil, false
}

func (g *Granary) Add(conn *websocket.Conn, user uint, lv int64) error {
	for _, wheat := range g.WheatList {
		if wheat.UserID == user {
			return fmt.Errorf("repeated addition of user(%d)", user)
		}
	}
	max := src.GetSetting().DefaultSetting.DefaultMaxPCU
	if g.PCU < max {
		g.WheatList = append(g.WheatList, &Wheat{
			Conn:          conn,
			UserID:        user,
			LatestVersion: lv,
		})
		g.PCU++
		return nil
	}
	return fmt.Errorf("there are too many connections to the same canvas, up to %d allowed", max)
}

func (g *Granary) Remove(user uint) bool {
	n := make([]*Wheat, 0)
	for _, wheat := range g.WheatList {
		if wheat.UserID != user {
			n = append(n, wheat)
		}
	}
	if len(n) == g.PCU {
		return false
	}
	g.PCU--
	g.WheatList = n
	return true
}

func (g *Granary) NotifyAll(pong pingpong.PongMessage) {
	data, _ := json.Marshal(pong)
	for _, wheat := range g.WheatList {
		err := wheat.Notify(data)
		if err != nil {
			log.Printf("fail to notify user(%d), data: %v, err: %v", wheat.UserID, pong, err)
		}
	}
}

func (g *Granary) NotifyOne(user uint, pong pingpong.PongMessage) bool {
	data, _ := json.Marshal(pong)
	wheat, ok := g.Find(user)
	if !ok {
		return false
	}
	err := wheat.Notify(data)
	if err != nil {
		log.Printf("fail to notify user(%d), data: %v, err: %v", user, pong, err)
		return false
	}
	return true
}

type connectWarehouse struct {
	depot sync.Map
	m     sync.Mutex
}

func (c *connectWarehouse) Load(canvasID uint) (*Granary, bool) {
	g, ok := c.depot.Load(canvasID)
	if !ok {
		return nil, false
	}
	granary := g.(*Granary)
	return granary, true
}

// HasUser 返回 user 是不是 canvas 的协作者
func (c *connectWarehouse) HasUser(canvasID, userID uint) bool {
	g, ok := c.depot.Load(canvasID)
	if !ok {
		return false
	}
	granary := g.(*Granary)
	_, ok = granary.Find(userID)
	return ok
}

// AddConnect 插入一个关于 canvasID 的 WS 连接
func (c *connectWarehouse) AddConnect(canvasID, userID uint, lv int64, conn *websocket.Conn) error {
	granary, stored := c.depot.LoadOrStore(canvasID, &Granary{
		WheatList: []*Wheat{{
			Conn:          conn,
			UserID:        userID,
			LatestVersion: lv,
		}},
		CanvasID:       canvasID,
		PCU:            1,
		RootRejoinSign: make(chan struct{}, 1),
	})
	if stored {
		g := granary.(*Granary)
		return g.Add(conn, userID, lv)
	}
	return nil
}

// RemoveConnect 从 canvasID 的连接者中删除 userID 对应的连接
func (c *connectWarehouse) RemoveConnect(canvasID, userID uint) error {
	c.m.Lock()
	defer c.m.Unlock()
	g, ok := c.depot.Load(canvasID)
	if !ok {
		return fmt.Errorf("[global/wspool] key(%d) not find when remove connect", canvasID)
	}
	granary := g.(*Granary)
	if !granary.Remove(userID) {
		return fmt.Errorf("[global/wspool] user(%d) not find when remove connect", userID)
	}
	return nil
}

func (c *connectWarehouse) RemoveCanvas(canvasID uint) {
	c.depot.Delete(canvasID)
}

// GetCanvasPCU 返回 canvas 的协作者数量
func (c *connectWarehouse) GetCanvasPCU(canvasID uint) int {
	g, ok := c.Load(canvasID)
	if !ok {
		return 0
	}
	return g.PCU
}

var webSocketPool *connectWarehouse
var wspLock sync.Mutex

func GetWSPool() *connectWarehouse {
	if webSocketPool != nil {
		return webSocketPool
	}
	wspLock.Lock()
	if webSocketPool == nil {
		webSocketPool = &connectWarehouse{depot: sync.Map{}}
	}
	wspLock.Unlock()
	return webSocketPool
}
