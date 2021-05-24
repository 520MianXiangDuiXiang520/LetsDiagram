package pingpong

import (
    `lets_diagram/src/core/access`
)

type PingType int

const (
    PingTypeVersionControl    PingType = iota + 1 // 版本控制信息
    PingTypeHeartbeat                             // 心跳包
    PingTypePermissionControl                     // 权限控制信息
)

type PingPermissionControlType int

const (
    PingPermissionApplication PingPermissionControlType = iota + 1 // 只读用户请求获得写权限
    PingPermissionAllowed                                          // 管理员允许某个只读用户的请求
    PingPermissionDenied                                           // 管理员拒绝某个只读用户的请求
    PingPermissionSet                                              // 管理员设置某个用户的权限
)

// PingPermissionData 当 ping 的类型为 PingTypePermissionControl 时的 data
type PingPermissionData struct {
    Type          PingPermissionControlType     `json:"type"`
    User          uint                          `json:"user"` // 操作对象
    NewPermission access.CanvasJurisdictionMark `json:"new_permission"`
}

type PingMessage struct {
    Type PingType `json:"type"`
    Data Message  `json:"data"`
}

type PingPermissionMsg struct {
    Type PingType           `json:"type"`
    Data PingPermissionData `json:"data"`
}

type PingVersionControlData struct {
    BaseVersion int64       `json:"base_version"`
    Diffs       interface{} `json:"diffs"`
}

type PingVersionControlMsg struct {
    Type PingType               `json:"type"`
    Data PingVersionControlData `json:"data"`
}






