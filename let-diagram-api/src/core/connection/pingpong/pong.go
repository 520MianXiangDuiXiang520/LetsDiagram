package pingpong

import (
    `lets_diagram/src/core/access`
    `lets_diagram/src/models`
    `lets_diagram/src/models/nottable`
)

type PongType int

const (
    PongTypeVersionControl        PongType = iota + 1 // 版本控制信息
    PongTypePermissionControl                         // 权限控制信息
    PongTypeHeartbeatControl                          // 第一次同步
    PongTypeRootCloseStopTheWorld                     // 画布创建者已经离开协作，所有人禁止写，等待创建者重新加入
    PongTypeSimpleNotify                              // 负载一些简单的通知信息
)

type PongSimpleNotifyType int

const (
    PongSimpleNotifyUserAdd PongSimpleNotifyType = iota + 1  // 有用户加入协作
    PongSimpleNotifyUserOut                                  // 有用户离开协作
)

type PongSimpleNotifyUserAddData struct {
    Type PongSimpleNotifyType `json:"type"`
    User *nottable.SimpleUser `json:"user"`
}

type PongSimpleNotifyUserOutData PongSimpleNotifyUserAddData

type PongRootCloseSTWType int

const (
    PongRootCloseSTWCrash = iota + 1  // 所有人停止协作
    PongRootCloseSTWRecovery          // 恢复协作
)

type PongPermissionControlType int

const (
    PongPermissionApplication PongPermissionControlType = iota + 1 // 有人向你申请
    PongPermissionAllowed                                          // 管理员通过了你的申请
    PongPermissionDenied                                           // 管理员拒绝了你的申请
    PongPermissionSet                                              // 有人修改了你的权限
    PongPermissionKickOut                                          // 管理员将你踢出了协作
)

type PongSTWData struct {
    Type PongRootCloseSTWType `json:"type"`
}

// PongPermissionData 当 pong 类型为 PongTypePermissionControl 时的 Data
type PongPermissionData struct {
    Type          PongPermissionControlType     `json:"type"` // 有人向你申请？有人拒绝/通过了你的申请？
    User          *nottable.SimpleUser            `json:"user"` // 谁向你申请？谁通过/拒绝了你的申请
    NewPermission access.CanvasJurisdictionMark `json:"new_permission"`
}

type PongMessage struct {
    Type PongType `json:"type"`
    Data Message  `json:"data"`
}

type PongPermissionMsg struct {
    Type PongType           `json:"type"`
    Data PongPermissionData `json:"data"`
}

type PongVersionControlData struct {
    Version     int64        `json:"version"`
    Diffs       interface{}  `json:"diffs"`
    PCU         int          `json:"pcu"` // 在线人数
    User        *models.User `json:"user"`
    ISWriter      bool         `json:"is_writer"`
    BaseVersion int64        `json:"base_version"`
}

type PongVersionControlMsg struct {
    Type PongType                 `json:"type"`
    Data []PongVersionControlData `json:"data"`
}

type PongHeartbeatData PongVersionControlData
type PongHeartbeatMsg struct {
    Type PongType            `json:"type"`
    Data []PongHeartbeatData `json:"data"`
}
