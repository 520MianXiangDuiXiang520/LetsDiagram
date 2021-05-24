package access

import (
    "sync"
)

// DAC 自主访问控制

type CanvasJurisdictionMark uint8

const (

	// CanvasJurisdictionMarkNotAnyAuthorized 用户无权参与该 canvas 的协作
	// 此项权限不会主动授予，主要用于获取用户权限时的标识。
	CanvasJurisdictionMarkNotAnyAuthorized CanvasJurisdictionMark = iota + 1

	// CanvasJurisdictionMarkReadOnly 用户通过分享加入协作后的默认权限，此权限的
	// 用户只能看到协作绘图的结果，但不能对 canvas 进行修改，
	// 但可以向拥有 CanvasJurisdictionMarkManager 权限的用户申请 write 权限
	CanvasJurisdictionMarkReadOnly

	// CanvasJurisdictionMarkWrite 拥有该权限的用户可以对 canvas 内容进行修改
	CanvasJurisdictionMarkWrite

	// CanvasJurisdictionMarkManager 持有该权限的用户可以授予或剥夺低权限用户的
	// CanvasJurisdictionMarkWrite 和 CanvasJurisdictionMarkReadOnly 权限
	CanvasJurisdictionMarkManager

	// CanvasJurisdictionMarkRoot 该权限由画布创建者持有，不被转移或删除，拥有该
	// 权限的用户可以授予或剥夺其他用户的其他权限。
	CanvasJurisdictionMarkRoot
)

const RootSetPermission = uint(0)

type accessControlRule map[uint]struct{}

func (a accessControlRule) hasUser(u uint) bool {
	_, ok := a[u]
	return ok
}

func (a accessControlRule) addUser(u uint) {
	a[u] = struct{}{}
}
func (a accessControlRule) removeUser(u uint) {
	delete(a, u)
}

func (a accessControlRule) length() int {
    return len(a)
}

// 每一个 canvas 对应一个 accessControlRow, 用来记录
// 与该 canvas 相关的访问控制信息
type accessControlRow struct {
	rooter  uint
	manager accessControlRule
	writer  accessControlRule
	reader  accessControlRule
	m       sync.Mutex
	stw     bool
}

type accessControlTable struct {
	depot sync.Map
}

func (act *accessControlTable) Load(canvasID uint) (*accessControlRow, bool) {
    v, ok := act.depot.Load(canvasID)
    if !ok {
        return nil, false
    }
    acr := v.(*accessControlRow)
    return acr, true
}

func (act *accessControlTable) authentication(baseMark, toMark CanvasJurisdictionMark) bool {
	if baseMark == CanvasJurisdictionMarkRoot {
		return true
	}
	if baseMark == CanvasJurisdictionMarkManager {
		return toMark == CanvasJurisdictionMarkWrite || toMark == CanvasJurisdictionMarkReadOnly
	}
	return false
}

// SetPermission 将 per 权限授予正在参与 canvas 协作的 user。
// per 只能是 CanvasJurisdictionMarkReadOnly， CanvasJurisdictionMarkWrite 或
// CanvasJurisdictionMarkManager , 如果该用户当前拥有其他权限，其他权限会被撤销。
func (act *accessControlTable) SetPermission(youID, canvasID, userID uint, per CanvasJurisdictionMark) {
	if youID != 0 && !act.authentication(act.GetPermission(canvasID, youID), act.GetPermission(canvasID, userID)) {
		return
	}
	acr, ok := act.Load(canvasID)
	if !ok {
        return
    }
	acr.m.Lock()
	defer acr.m.Unlock()
	switch per {
	case CanvasJurisdictionMarkReadOnly:
		if acr.writer.hasUser(userID) {
			acr.writer.removeUser(userID)
		}
		if acr.manager.hasUser(userID) {
			acr.manager.removeUser(userID)
		}
		acr.reader.addUser(userID)
	case CanvasJurisdictionMarkWrite:
		if acr.reader.hasUser(userID) {
			acr.reader.removeUser(userID)
		}
		if acr.manager.hasUser(userID) {
			acr.manager.removeUser(userID)
		}
		acr.writer.addUser(userID)
	case CanvasJurisdictionMarkManager:
		if acr.writer.hasUser(userID) {
			acr.writer.removeUser(userID)
		}
		if acr.reader.hasUser(userID) {
			acr.reader.removeUser(userID)
		}
		acr.manager.addUser(userID)
	}
}

// InitPermission 初始化 canvas 的访问控制表，并将 user 设置为该 canvas 的 root 用户
func (act *accessControlTable) InitPermission(canvasID, userID uint) {
	value, stored := act.depot.LoadOrStore(canvasID, &accessControlRow{
		rooter:  userID,
		manager: accessControlRule{},
		writer:  accessControlRule{},
		reader:  accessControlRule{},
		m:       sync.Mutex{},
	})
	if stored {
	    ac := value.(*accessControlRow)
	    ac.rooter = userID
    }
}

// GetPermission 返回 user 在 canvas 上的权限，默认 无任何权限
func (act *accessControlTable) GetPermission(canvasID, userID uint) CanvasJurisdictionMark {
    acr, ok := act.Load(canvasID)
    if !ok {
        return CanvasJurisdictionMarkNotAnyAuthorized
    }
	if acr.stw {
	    return CanvasJurisdictionMarkNotAnyAuthorized
    }
	if acr.rooter == userID {
		return CanvasJurisdictionMarkRoot
	}
	if acr.manager.hasUser(userID) {
		return CanvasJurisdictionMarkManager
	}
	if acr.writer.hasUser(userID) {
		return CanvasJurisdictionMarkWrite
	}
	if acr.reader.hasUser(userID) {
		return CanvasJurisdictionMarkReadOnly
	}
	return CanvasJurisdictionMarkNotAnyAuthorized
}

// GetRoot 返回 canvas 中，拥有写权限的 userID, 如果 canvas 中无人拥有写权限，返回 (0, false)
func (act *accessControlTable) GetRoot(canvasID uint) (uint, bool) {
    acr, ok := act.Load(canvasID)
    if !ok || acr.rooter == 0 {
        return 0, false
    }
	return acr.rooter, true
}

// GetBosses 获取当前 canvas 的所有管理者（root or manager）ID
func (act *accessControlTable) GetBosses(canvasID uint) []uint {
    acr, ok := act.Load(canvasID)
    if !ok {
        return nil
    }
	result := make([]uint, len(acr.manager)+1)
	result[0] = acr.rooter
	idx := 0
	for id, _ := range acr.manager {
		result[idx+1] = id
		idx++
	}
	return result
}

// DriveOut 删除一个用户关于 Canvas 的全部权限
func (act *accessControlTable) DriveOut(canvasID, userID uint) {
    acr, ok := act.Load(canvasID)
    if !ok {
        return
    }
	acr.m.Lock()
	if acr.rooter == userID {
	    acr.rooter = 0
    }
	acr.reader.removeUser(userID)
	acr.writer.removeUser(userID)
	acr.manager.removeUser(userID)
	if acr.rooter == 0 && acr.writer.length() == 0 &&
	    acr.reader.length() == 0 && acr.manager.length() == 0 {
        act.depot.Delete(canvasID)
    }
	acr.m.Unlock()
}

// GetList 返回该 canvas 的所有协作者
func (act *accessControlTable) GetList(canvasID uint) map[CanvasJurisdictionMark][]uint {
    acr, ok := act.Load(canvasID)
    if !ok {
        return nil
    }
	result := make(map[CanvasJurisdictionMark][]uint)
	result[CanvasJurisdictionMarkRoot] = []uint{acr.rooter}

	managers := make([]uint, 0)
	for id, _ := range acr.manager {
		managers = append(managers, id)
	}
	result[CanvasJurisdictionMarkManager] = managers

	writers := make([]uint, 0)
	for id, _ := range acr.writer {
		writers = append(writers, id)
	}
	result[CanvasJurisdictionMarkWrite] = writers

	readers := make([]uint, 0)
	for id, _ := range acr.reader {
		readers = append(readers, id)
	}
	result[CanvasJurisdictionMarkReadOnly] = readers
	return result
}

func (act *accessControlTable) STW(canvasID uint) {
    acr, ok := act.Load(canvasID)
    if !ok {
        return
    }
    acr.m.Lock()
    acr.stw = true
    acr.m.Unlock()
}

func (act *accessControlTable) UnSTW(canvasID uint) {
    acr, ok := act.Load(canvasID)
    if !ok {
        return
    }
    acr.m.Lock()
    acr.stw = false
    acr.m.Unlock()
}

func (act *accessControlTable) IsSTW(canvasID uint) bool {
    acr, ok := act.Load(canvasID)
    return ok && acr.stw
}

func (act *accessControlTable) RemoveCanvas(canvasID uint) {
    act.depot.Delete(canvasID)
}

var act *accessControlTable
var actMutex = sync.Mutex{}

func GetAccessControlTable() *accessControlTable {
    if act == nil {
        actMutex.Lock()
        if act == nil {
            act = &accessControlTable{depot: sync.Map{}}
        }
        actMutex.Unlock()
    }
    return act
}

// CanSetPermission 返回 user 在 canvas 上是否具有管理员（root or manager）权限
func CanSetPermission(canvasID, userID uint) bool {
	return GetAccessControlTable().GetPermission(canvasID, userID) >= CanvasJurisdictionMarkManager
}

// CanWrite 返回 user 是否拥有在 canvas 上的写权限
func CanWrite(canvasID, userID uint) bool {
	return GetAccessControlTable().GetPermission(canvasID, userID) > CanvasJurisdictionMarkReadOnly
}
