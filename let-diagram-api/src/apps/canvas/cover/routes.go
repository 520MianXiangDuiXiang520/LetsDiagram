package cover

import (
    `github.com/520MianXiangDuiXiang520/ginUtils`
    `github.com/gin-gonic/gin`
)

func Routes(g *gin.RouterGroup) {
    g.POST("update/", coverUpdateHandler()...)
    g.POST("get/",    coverGetHandler()...)
}

func coverGetHandler() []gin.HandlerFunc {
    return []gin.HandlerFunc{
        ginUtils.Handler(GetCheck, GetLogic, GetRequestFields{}),
    }
}

func coverUpdateHandler() []gin.HandlerFunc {
    return []gin.HandlerFunc{
        ginUtils.Handler(UpdateCheck, UpdateLogic, UpdateRequestFields{}),
    }
}
