package versionlink

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"lets_diagram/src/core/access"
	"lets_diagram/src/models"
	"lets_diagram/src/utils"
	"log"
	"sync"
	"time"
)

type DiffNode struct {
	Value       interface{}  `json:"value"`   // 该版本修改了哪些值
	UserID      uint         `json:"user_id"` // 该版本的创建者
	User        *models.User `json:"user"`
	BaseVersion int64        `json:"base_version"`
	Version     int64        `json:"version"` // 版本号
	Time        int64        `json:"time"`    // 版本创建时间
}

func CreateDiffNode(value interface{}, user *models.User, baseVersion int64) *DiffNode {
	return &DiffNode{Value: value, UserID: user.ID, BaseVersion: baseVersion, User: user,
		Time: time.Now().UnixNano()}
}

type diffLinkNode struct {
	DiffNode
	next *diffLinkNode
	prev *diffLinkNode
}

// DiffLink 表示某个 Canvas 对应的 diff 版本链
// 每次客户端对 Canvas 的修改都会产生一个新的版本块，并插入到版本链的头部
// 版本链的头部表示的是最新的版本，存储模块定时从版本链尾部摘取版本块去做持久化
type DiffLink struct {
	head           *diffLinkNode
	tail           *diffLinkNode
	Length         int
	LatestVersion  int64
	bodyMutex      sync.RWMutex // 锁定整个版本链
	headMutex      sync.RWMutex // 锁定版本链头部
	tailMutex      sync.RWMutex // 锁定版本链尾部
	EnduranceMutex sync.Mutex   // 保证持久化任务一致性, 删除 canvas 时，业务 goroutine 会与持久化 canvas 竞争该锁
	skipEndurance  bool
}

// NewVersion 将一个最新的版本块插入到版本链的首部
func (dl *DiffLink) NewVersion(node *DiffNode) {
	// 丢弃对同一版本数据的多次修改，保证同时只有一个用户修改上一版本的数据
	if node.BaseVersion != 0 && node.BaseVersion != dl.LatestVersion {
		log.Printf("丢弃对同一版本数据的多次修改，保证同时只有一个用户修改上一版本的数据, base(%d), last(%d) \n",
			node.BaseVersion, dl.LatestVersion)
		return
	}

	n := &diffLinkNode{
		prev: nil,
		next: dl.head,
	}
	n.Version = dl.LatestVersion + 1
	n.Value = node.Value
	n.BaseVersion = node.BaseVersion
	n.UserID = node.UserID
	n.Time = node.Time
	n.User = node.User

	// 检查空链表
	if dl.head == nil && dl.tail == nil {
		dl.headMutex.Lock()
		dl.tailMutex.Lock()
		if dl.head == nil && dl.tail == nil {
			dl.head = n
			dl.tail = n
			dl.Length++
			dl.LatestVersion++
		}
		dl.tailMutex.Unlock()
		dl.headMutex.Unlock()
	} else {
		dl.headMutex.Lock()
		dl.head.prev = n
		dl.head = n
		dl.Length++
		dl.LatestVersion++
		dl.headMutex.Unlock()
	}
}

func (dl *DiffLink) String() string {
	buffer := bytes.NewBufferString("")
	cur := dl.head
	for ; cur != nil; cur = cur.next {
		buffer.WriteString(fmt.Sprintf("{version: %d, userID: %d, baseVersion: %d} -> ",
			cur.Version, cur.UserID, cur.BaseVersion))
	}
	return buffer.String()
}

func (dl *DiffLink) RemoveN(n int) int {
	dl.bodyMutex.Lock()
	defer dl.bodyMutex.Unlock()

	dl.tailMutex.Lock()
	defer dl.tailMutex.Unlock()
	if dl.Length <= n {
		dl.headMutex.Lock()
		res := dl.Length
		dl.head = nil
		dl.tail = nil
		dl.Length = 0
		dl.headMutex.Unlock()
		return res
	}
	var index int
	for dl.tail != nil && index < n {
		cur := dl.tail
		dl.tail = dl.tail.prev
		if dl.tail != nil {
			dl.tail.next = nil
		} else {
			dl.head = nil
		}
		cur.prev = nil
		cur = nil
		dl.Length--
		index++
	}
	return index
}

// GetLastN 返回版本链尾部的 n 个节点，第二个参数表示实际返回的节点数
func (dl *DiffLink) GetLastN(n int) ([]*DiffNode, int) {
	dl.tailMutex.RLock()
	defer dl.tailMutex.RUnlock()
	if dl.tail == nil {
		return nil, 0
	}
	result := make([]*DiffNode, n)
	index := 0
	cur := dl.tail
	for cur != nil && index < n {
		result[index] = &cur.DiffNode
		index++
		cur = cur.prev
	}
	return result, index
}

// Snapshot  stop the world and return a deep copy of this version chain
func (dl *DiffLink) Snapshot() (*DiffLink, error) {
	dl.headMutex.Lock()
	dl.tailMutex.Lock()
	defer func() {
		dl.tailMutex.Unlock()
		dl.headMutex.Unlock()
	}()
	cp := new(DiffLink)
	err := utils.DeepCopy(cp, dl)
	if err != nil {
		return nil, errors.Wrap(err, "snapshot can not to deep copy version link")
	}
	return cp, nil
}

// UnPushVersions 返回版本链中所有大于 version 的版本块（需要同步给客户端的所有版本）
// 返回的版本块是按版本号从新到旧排列的，客户端根据 diff 还原时，需要逆序遍历还原
func (dl *DiffLink) UnPushVersions(version int64) []*DiffNode {
	dl.headMutex.RLock()
	defer dl.headMutex.RUnlock()
	res := make([]*DiffNode, 0)
	cur := dl.head
	for cur != nil && cur.Version > version {
		res = append(res, &DiffNode{
			Value:       cur.Value,
			Version:     cur.Version,
			UserID:      cur.UserID,
			User:        cur.User,
			BaseVersion: cur.BaseVersion,
		})
		cur = cur.next
	}
	return res
}

func (dl *DiffLink) SetSkipEndurance() {
	dl.bodyMutex.Lock()
	dl.skipEndurance = true
	dl.bodyMutex.Unlock()
}

func (dl *DiffLink) SetUnSkipEndurance() {
	dl.bodyMutex.Lock()
	dl.skipEndurance = false
	dl.bodyMutex.Unlock()
}

const MaxDWHSharkNumber = 4

type diffGroup struct {
	depot sync.Map
	m     sync.Mutex
	cron  *cron.Cron
}

// 持久化任务
func (d *diffGroup) enduranceTask() {
	// 读版本链
	d.depot.Range(func(k, v interface{}) bool {
		canvasID, versionLink := k.(uint), v.(*DiffLink)
		versionLink.EnduranceMutex.Lock()
		if !versionLink.skipEndurance {
			n := 30
			nodes, count := versionLink.GetLastN(n)
			if d != nil && (count == 0 || nodes == nil) {
				// 检查 canvas 是否关闭
				if _, ok := access.GetAccessControlTable().Load(canvasID); !ok {
					// 删除该 canvas 的版本链
					d.delete(canvasID)
				}
				versionLink.EnduranceMutex.Unlock()
				return true
			}
			if NewEndurance(nodes, count, canvasID) {
				versionLink.RemoveN(count)
			}
		}
		versionLink.EnduranceMutex.Unlock()
		return true
	})
}

// 第一次打开一个 Canvas 时，取出上一次持久化的版本号初始化该 canvas 对应的版本链
func (d *diffGroup) newLink(canvasID uint, lastVersion int64) {
	dl := &DiffLink{
		Length:        0,
		LatestVersion: lastVersion,
		skipEndurance: false,
	}
	d.depot.LoadOrStore(canvasID, dl)
}

func (d *diffGroup) load(canvasID uint) (*DiffLink, bool) {
	v, ok := d.depot.Load(canvasID)
	if !ok {
		return nil, false
	}
	return v.(*DiffLink), ok
}

func (d *diffGroup) store(canvasID uint, node *DiffNode) {
	link, ok := d.load(canvasID)
	if ok {
		link.NewVersion(node)
	}
}

func (d *diffGroup) delete(canvasID uint) {
	d.depot.Delete(canvasID)
}

// DiffWareHouse 用来保存所有 canvas 的 diff 链，以 canvasID 为 key
type DiffWareHouse struct {
	shark [MaxDWHSharkNumber]*diffGroup
}

func (d *DiffWareHouse) Load(canvasID uint) (*DiffLink, bool) {
	key := canvasID % MaxDWHSharkNumber
	return d.shark[key].load(canvasID)
}

func (d *DiffWareHouse) Store(canvasID uint, node *DiffNode) {
	key := canvasID % MaxDWHSharkNumber
	d.shark[key].store(canvasID, node)
}

func (d *DiffWareHouse) Delete(canvasID uint) {
	key := canvasID % MaxDWHSharkNumber
	d.shark[key].delete(canvasID)
}

// UnPushVersions 从 CanvasID 对应的版本链中找到所有版本大于 version 的节点并返回；
// 返回的节点按版本从大到小排列。
func (d *DiffWareHouse) UnPushVersions(canvasID uint, version int64) ([]*DiffNode, bool) {
	l, ok := d.Load(canvasID)
	if ok {
		return l.UnPushVersions(version), true
	}
	return nil, false
}

// AddNewVersion 为 canvasID 对应的版本链插入一个新版本节点；
// userID 表示该节点的创建者，diffs 表示该节点的值；
// 返回 false 表示 canvasID 对应的版本链不存在。
func (d *DiffWareHouse) AddNewVersion(canvasID uint, node *DiffNode) bool {
	l, ok := d.Load(canvasID)
	if !ok {
		return false
	}
	l.NewVersion(node)
	return true
}

func (d *DiffWareHouse) NewVersionsLink(canvasID uint, lastVersion int64) {
	key := canvasID % MaxDWHSharkNumber
	d.shark[key].newLink(canvasID, lastVersion)
}

var dwh *DiffWareHouse
var dwhLock sync.Mutex

func GetVersionsLink() *DiffWareHouse {
	if dwh != nil {
		return dwh
	}
	dwhLock.Lock()
	if dwh == nil {
		dwh = &DiffWareHouse{shark: [MaxDWHSharkNumber]*diffGroup{}}
		for i := 0; i < MaxDWHSharkNumber; i++ {
			dwh.shark[i] = &diffGroup{
				depot: sync.Map{},
				m:     sync.Mutex{},
				cron:  cron.New(cron.WithSeconds()),
			}
			_, _ = dwh.shark[i].cron.AddFunc("@every 10s", dwh.shark[i].enduranceTask)
			dwh.shark[i].cron.Start()
		}
	}
	dwhLock.Unlock()
	return dwh
}
