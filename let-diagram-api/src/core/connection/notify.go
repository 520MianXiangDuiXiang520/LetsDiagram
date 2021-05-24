package connection

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"lets_diagram/src/core/connection/pingpong"
	"lets_diagram/src/core/versionlink"
	"log"
)

// sendVersionPong 将版本链中所有版本号大于 wheat.LatestVersion 的节点推送给 wheat
func sendVersionPong(g *Granary, diffLink *versionlink.DiffLink, wheat *Wheat) {
	unPush := diffLink.UnPushVersions(wheat.LatestVersion)
	size := len(unPush)
	if size <= 0 {
		return
	}
	vcMsg := make([]pingpong.PongVersionControlData, size)
	for i := size - 1; i >= 0; i-- {
		up := unPush[i]
		pm := pingpong.PongVersionControlData{
			Version:     up.Version,
			Diffs:       up.Value,
			User:        up.User,
			PCU:         g.PCU,
			ISWriter:    up.User.ID != wheat.UserID,
			BaseVersion: up.BaseVersion,
		}
		vcMsg[size-i-1] = pm
	}
	pongMsg := pingpong.PongVersionControlMsg{
		Type: pingpong.PongTypeVersionControl,
		Data: vcMsg,
	}
	data, _ := json.Marshal(pongMsg)
	err := wheat.Notify(data)
	if err != nil {
		log.Printf("[global/notifyAll] failed to send to %d, error info: %v", wheat.UserID, err)
		return
	}
	wheat.LatestVersion = unPush[0].Version
}

func sendHeartbeatPong(g *Granary, diffLink *versionlink.DiffLink, wheat *Wheat) {
	unPush := diffLink.UnPushVersions(wheat.LatestVersion)
	size := len(unPush)
	if size <= 0 {
		return
	}
	vcMsg := make([]pingpong.PongHeartbeatData, size)
	for i := size - 1; i >= 0; i-- {
		up := unPush[i]
		pm := pingpong.PongHeartbeatData{
			Version:     up.Version,
			Diffs:       up.Value,
			PCU:         g.PCU,
			ISWriter:    true,
			BaseVersion: up.BaseVersion,
		}
		vcMsg[size-i-1] = pm
	}
	pongMsg := pingpong.PongHeartbeatMsg{
		Type: pingpong.PongTypeHeartbeatControl,
		Data: vcMsg,
	}
	data, _ := json.Marshal(pongMsg)
	err := wheat.Notify(data)
	if err != nil {
		log.Printf("[global/notifyAll] failed to send to %d, error info: %v", wheat.UserID, err)
		return
	}
	wheat.LatestVersion = unPush[0].Version
}

// VersionNotifyAll 将最新的版本信息推送给 granary 中的所有协作者
func VersionNotifyAll(canvasID uint) {
	g, ok := GetWSPool().Load(canvasID)
	if !ok {
		return
	}
	diffLink, ok := versionlink.GetVersionsLink().Load(g.CanvasID)
	if !ok {
		log.Printf("no diff link")
		return
	}
	for _, wheat := range g.WheatList {
		sendVersionPong(g, diffLink, wheat)
	}
}

// VersionNotifyOne 将最新的版本信息推送给特定的用户
func VersionNotifyOne(canvasID, userID uint) {
	g, ok := GetWSPool().Load(canvasID)
	if !ok {
		log.Printf("not find in connectWarehouse.VersionNotifyOne")
		return
	}
	diffLink, ok := versionlink.GetVersionsLink().Load(g.CanvasID)
	if !ok {
		log.Printf("no diff link")
		return
	}
	wheat, ok := g.Find(userID)
	if !ok {
		log.Printf("could not find the connection for this user(%d)", userID)
		return
	}
	sendHeartbeatPong(g, diffLink, wheat)
}

func SendPong(canvasID, user uint, pong pingpong.PongMessage) {
	g, ok := GetWSPool().Load(canvasID)
	if !ok {
		return
	}
	wheat, ok := g.Find(user)
	if !ok {
		return
	}
	pongMessage, err := json.Marshal(pong)
	if err != nil {
		return
	}
	_ = wheat.Conn.WriteMessage(websocket.TextMessage, pongMessage)
}

func SendAllPong(canvasID uint, pong pingpong.PongMessage) {
	g, ok := GetWSPool().Load(canvasID)
	if !ok {
		return
	}
	pongMessage, err := json.Marshal(pong)
	if err != nil {
		return
	}
	for _, wheat := range g.WheatList {
		_ = wheat.Notify(pongMessage)
	}
}

func STWPongNotifyAll(canvasID uint, t pingpong.PongRootCloseSTWType) {
	pong := pingpong.GetSTWPong(t)
	g, ok := GetWSPool().Load(canvasID)
	if !ok {
		return
	}
	g.NotifyAll(pong)
}

func STWPongNotifyOne(canvasID, userID uint, t pingpong.PongRootCloseSTWType) {
	pong := pingpong.GetSTWPong(t)
	g, ok := GetWSPool().Load(canvasID)
	if !ok {
		return
	}
	g.NotifyOne(userID, pong)
}
