package handler

import (
    "encoding/json"
    "github.com/gorilla/websocket"
    `lets_diagram/src/core/access`
    `lets_diagram/src/core/connection`
    `lets_diagram/src/core/connection/pingpong`
    "lets_diagram/src/core/versionlink"
    "lets_diagram/src/models"
    "log"
)

func VersionControlHandler(message []byte, cID uint, user *models.User, ws *websocket.Conn) {
	var ping pingpong.PingVersionControlMsg
	err := json.Unmarshal(message, &ping)
	if err != nil {
		log.Printf("[apps/canvas/server] versionControlHandler() Unmarshal fail: %v", err)
		return
	}
	// 如果当前用户拥有写权限，修改版本链
	if access.GetAccessControlTable().GetPermission(cID, user.ID) >= access.CanvasJurisdictionMarkWrite {
		versionlink.GetVersionsLink().AddNewVersion(cID,
			versionlink.CreateDiffNode(ping.Data.Diffs, user, ping.Data.BaseVersion))
        // 通知所有用户更新
        connection.VersionNotifyAll(cID)
	}
	
}
