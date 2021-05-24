package canvas

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"lets_diagram/src/core/connection"
	"lets_diagram/src/models"
	"reflect"
	"strconv"
)

// SaveConn 是一个 TransformToWS 的装饰器，用来保存一份与服务端建立了连接的 websocket conn
// 的引用到 global.WebSocketConnectionPool 中，该装饰器需要获取 userID, 所以在这之前需要
// 保证 ginCtx.Get("user") 咳哟获取到正确的结果。
func SaveConn(ws *websocket.Conn, ginCtx *gin.Context) bool {
	canvasID := ginCtx.Query("canvas_id")
	// TODO: 加密
	cID, err := strconv.Atoi(canvasID)
	if err != nil {
		return false
	}
	u, ok := ginCtx.Get("user")
	if !ok {
		return false
	}
	user := u.(*models.User)

	version := ginCtx.Query("version")
	lv, err := strconv.Atoi(version)
	if err != nil {
		return false
	}
	if err := connection.GetWSPool().AddConnect(uint(cID), user.ID, int64(lv), ws); err != nil {
		return false
	}
	return true
}

func BindData(ws *websocket.Conn, ginCtx *gin.Context) bool {
	t := reflect.TypeOf(PaintWSRequestFields{})
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("json")
		v := ginCtx.Query(tag)
		ginCtx.Set(tag, v)
	}
	return true
}
