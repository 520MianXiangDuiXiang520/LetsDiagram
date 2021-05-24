package handler

import (
    "encoding/json"
    "github.com/gorilla/websocket"
    `lets_diagram/src/core/access`
    `lets_diagram/src/core/connection`
    `lets_diagram/src/core/connection/pingpong`
    "lets_diagram/src/models"
    "log"
)

func PermissionHandler(message []byte, cID uint, user *models.User, ws *websocket.Conn) {
	var ping pingpong.PingPermissionMsg
	err := json.Unmarshal(message, &ping)
	if err != nil {
		log.Printf("[apps/canvas/server] permissionHandler() Unmarshal fail: %v", err)
		return
	}
	pcMessage := ping.Data
	switch pcMessage.Type {
	case pingpong.PingPermissionApplication:
		// 只有只读用户能请求
		if access.GetAccessControlTable().GetPermission(cID, user.ID) != access.CanvasJurisdictionMarkReadOnly {
			log.Printf("%d 的用户申请写权限失败！", access.GetAccessControlTable().GetPermission(cID, user.ID))
			return
		}
		// 通知所有管理员
		connection.SendPong(cID, pcMessage.User, pingpong.GetPermissionPong(user,
            pingpong.PongPermissionApplication,
            access.CanvasJurisdictionMarkWrite))
	case pingpong.PingPermissionAllowed:
		u := pcMessage.User // 同意谁
		if access.CanWrite(cID, u) {
			// 如果用户已经有了写权限，说明其他管理员已同意，暂时忽略
			return
		}
		access.GetAccessControlTable().SetPermission(user.ID, cID, u, access.CanvasJurisdictionMarkWrite)
		connection.SendPong(cID, u, pingpong.GetPermissionPong(user,
            pingpong.PongPermissionAllowed, access.CanvasJurisdictionMarkWrite))
	case pingpong.PingPermissionDenied:
		if !access.CanSetPermission(cID, user.ID) {
			return
		}
		connection.SendPong(cID, pcMessage.User, pingpong.GetPermissionPong(user,
            pingpong.PongPermissionDenied, access.CanvasJurisdictionMarkWrite))
	case pingpong.PingPermissionSet:
		if !access.CanSetPermission(cID, user.ID) {
			log.Printf("无权修改")
			return
		}
		if access.GetAccessControlTable().GetPermission(cID, pcMessage.User) == pcMessage.NewPermission {
			return
		}
		switch pcMessage.NewPermission {
		case access.CanvasJurisdictionMarkNotAnyAuthorized:
			// 删除 connection 和 权限记录
			connection.SendPong(cID, pcMessage.User, pingpong.GetPermissionPong(user,
                pingpong.PongPermissionKickOut, access.CanvasJurisdictionMarkNotAnyAuthorized))
			access.GetAccessControlTable().DriveOut(cID, pcMessage.User)
			_ = connection.GetWSPool().RemoveConnect(cID, pcMessage.User)
		case access.CanvasJurisdictionMarkWrite:
			access.GetAccessControlTable().SetPermission(user.ID, cID, pcMessage.User, access.CanvasJurisdictionMarkWrite)
			connection.SendPong(cID, pcMessage.User, pingpong.GetPermissionPong(user,
                pingpong.PongPermissionSet, access.CanvasJurisdictionMarkWrite))
		case access.CanvasJurisdictionMarkReadOnly:
			access.GetAccessControlTable().SetPermission(user.ID, cID, pcMessage.User, access.CanvasJurisdictionMarkReadOnly)
			connection.SendPong(cID, pcMessage.User, pingpong.GetPermissionPong(user,
                pingpong.PongPermissionSet, access.CanvasJurisdictionMarkReadOnly))
		case access.CanvasJurisdictionMarkManager:
			if access.GetAccessControlTable().GetPermission(cID, user.ID) != access.CanvasJurisdictionMarkRoot {
				return
			}
			access.GetAccessControlTable().SetPermission(user.ID, cID, pcMessage.User, access.CanvasJurisdictionMarkManager)
			connection.SendPong(cID, pcMessage.User, pingpong.GetPermissionPong(user,
                pingpong.PongPermissionSet, access.CanvasJurisdictionMarkManager))
		}
	}
}
