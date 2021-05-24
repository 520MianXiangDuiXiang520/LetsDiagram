package user

import (
	"github.com/520MianXiangDuiXiang520/ginUtils"
	"github.com/gin-gonic/gin"
)

func Routes(g *gin.RouterGroup) {
	g.POST("/login",
		ginUtils.Handler(LoginCheck, LoginLogic, LoginRequestFields{}))
	g.POST("/register",
		ginUtils.Handler(RegisterCheck, RegisterLogic, RegisterRequestFields{}))
}
