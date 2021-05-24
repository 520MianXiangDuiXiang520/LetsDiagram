package main

import (
    "github.com/520MianXiangDuiXiang520/ginUtils"
    "github.com/520MianXiangDuiXiang520/ginUtils/middleware"
    "github.com/gin-gonic/gin"
    `lets_diagram/src`
    "lets_diagram/src/apps/canvas"
    "lets_diagram/src/apps/user"
)

func rootRoutes(engine *gin.Engine) {
	engine.Use(middleware.CorsHandler(src.GetSetting().CORSSetting.AllowedList))
	ginUtils.URLPatterns(engine, "/user", user.Routes)
	ginUtils.URLPatterns(engine, "/canvas", canvas.Routes)
}
