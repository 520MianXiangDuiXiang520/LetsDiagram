package canvas

import (
	"github.com/520MianXiangDuiXiang520/ginUtils"
	"github.com/520MianXiangDuiXiang520/ginUtils/middleware"
	"lets_diagram/src/apps/canvas/cover"
	middleware2 "lets_diagram/src/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(g *gin.RouterGroup) {
	g.POST("new/", newCanvasHandlers()...)
	g.POST("open/", openCanvasHandlers()...)
	g.POST("all/", allCanvasHandlers()...)
	g.POST("cooperate/", cooperateHandler()...)
	g.POST("check_cooperate/", checkCooperateHandler()...)
	g.POST("join/", joinHandler()...)
	g.POST("collaborators/", collaboratorsHandler()...)
	g.POST("fork/", forkHandler()...)
	g.POST("delete/", deleteCanvasHandler()...)
	g.POST("rename/", renameCanvasHandler()...)
	g.GET("paint/", paintCanvasHandlers()...)
	ginUtils.URLPatterns(g, "cover/", cover.Routes,
		middleware.Auth(middleware2.TokenAuth),
		middleware.Permiter(middleware2.ISCreator),
	)
}

func renameCanvasHandler() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.Auth(middleware2.TokenAuth),
		ginUtils.Handler(RenameCheck, RenameLogic, RenameRequestFields{}),
	}
}

func deleteCanvasHandler() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.Auth(middleware2.TokenAuth),
		middleware.Permiter(middleware2.ISCreator),
		ginUtils.Handler(DeleteCheck, DeleteLogic, DeleteRequestFields{}),
	}
}

func forkHandler() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.Auth(middleware2.TokenAuth),
		ginUtils.Handler(ForkCheck, ForkLogic, ForkRequestFields{}),
	}
}

func collaboratorsHandler() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.Auth(middleware2.TokenAuth),
		// middleware.Permiter(middleware2.ISCreator),
		ginUtils.Handler(CollaboratorsCheck, CollaboratorsLogic, CollaboratorsRequestFields{}),
	}
}

func checkCooperateHandler() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.Auth(middleware2.TokenAuth),
		ginUtils.Handler(CheckCooperateCheck, CheckCooperateLogic, CheckCooperateRequestFields{}),
	}
}

func joinHandler() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.Auth(middleware2.TokenAuth),
		ginUtils.Handler(JoinCheck, JoinLogic, JoinRequestFields{}),
	}
}

func cooperateHandler() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.Auth(middleware2.TokenAuth),
		middleware.Permiter(middleware2.ISCreator),
		ginUtils.Handler(CooperateCheck, CooperateLogic, CooperateRequestField{}),
	}
}

func allCanvasHandlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.Auth(middleware2.TokenAuth),
		ginUtils.Handler(AllCheck, AllLogic, AllCanvasRequestField{}),
	}
}

func openCanvasHandlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.Auth(middleware2.TokenAuth),
		middleware.Permiter(middleware2.ISCreator),
		ginUtils.Handler(OpenCheck, OpenLogic, OpenCanvasRequestField{}),
	}
}

func newCanvasHandlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.Auth(middleware2.TokenAuth),
		ginUtils.Handler(NewCheck, NewLogic, NewCanvasRequestField{}),
	}
}

func paintCanvasHandlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.Auth(middleware2.WebSocketAuth),
		TransformToWS(PaintLogic, nil, SaveConn),
	}
}
