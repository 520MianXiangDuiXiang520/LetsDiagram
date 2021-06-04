package middleware

import (
	"github.com/520MianXiangDuiXiang520/ginUtils/middleware"
	"github.com/gin-gonic/gin"
	"lets_diagram/src/dao"
)

func TokenAuth(context *gin.Context) (middleware.UserBase, bool) {
	var token string
	token, err := context.Cookie("SESSIONID")
	if err != nil {
		token = context.Request.Header.Get("token")
		if token == "" {
			return nil, false
		}
	}
	user, ok := dao.GetUserByToken(token)
	if !ok {
		return nil, false
	}
	return user, true
}

// WebSocketAuth 用于 WebSocket 认证, token 通过 URL 传递
func WebSocketAuth(ctx *gin.Context) (middleware.UserBase, bool) {
	token := ctx.Query("token")
	if token == "" {
		return nil, false
	}
	user, ok := dao.GetUserByToken(token)
	if !ok {
		return nil, false
	}
	return user, true
}
