package handler

import (
    "github.com/gorilla/websocket"
    `lets_diagram/src/core/access`
    `lets_diagram/src/core/connection`
    `lets_diagram/src/core/connection/pingpong`
    "lets_diagram/src/models"
)

func HeartBeatHandler(message []byte, cID uint, user *models.User, ws *websocket.Conn) {
    if access.GetAccessControlTable().IsSTW(cID) {
        connection.STWPongNotifyOne(cID, user.ID, pingpong.PongRootCloseSTWCrash)
    }
	connection.VersionNotifyOne(cID, user.ID)
}
