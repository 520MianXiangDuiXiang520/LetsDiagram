
export const PingTypeVersionControl = 1;                 // 版本控制
export const PingTypeHeartbeat = 2;                      // 心跳
export const PingTypePermissionControl = 3;              // 权限管理

export const PongTypeVersionContral = 1;                 // 版本控制
export const PongTypePermissionControl = 2;              // 权限管理
export const PongTypeHeartbeatControl = 3;               // 心跳
export const PongTypeRootCloseStopTheWorld = 4;          // 画布创建者离开，所有人禁止写，开启倒计时
export const PongTypeSimpleNotify = 5;                   // 负载一些简单的通知信息

export const PongSimpleNotifyUserAdd = 1;                // 由用户加入协作
export const PongSimpleNotifyUserOut = 2;                // 由用户退出协作

export const PongRootCloseSTWCrash = 1;                  // stw
export const PongRootCloseSTWRecovery = 2;               // 恢复

export const PongPermissionApplication = 1;              // 有人向你申请
export const PongPermissionAllowed = 2;                  // writer 通过了你的申请
export const PongPermissionDenied = 3;                   // writer 拒绝了你的申请
export const PongPermissionSet = 4;                      // 管理员设置权限
export const PongPermissionKickOut = 5;                  // 踢人

export const PingPermissionApplication = 1;              // 申请写
export const PingPermissionAllowed = 2;                  // 允许写
export const PingPermissionDenied = 3;                   // 拒绝写
export const PingPermissionSet = 4;                      // 设置权限

export const CanvasJurisdictionMarkNotAnyAuthorized = 1  // 无权
export const CanvasJurisdictionMarkReadOnly = 2          // 只读
export const CanvasJurisdictionMarkWrite = 3             // 可写
export const CanvasJurisdictionMarkManager = 4           // 管理员
export const CanvasJurisdictionMarkRoot = 5              // 主持人
