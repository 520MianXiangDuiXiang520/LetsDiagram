package cover

import (
	"github.com/520MianXiangDuiXiang520/ginUtils"
	"github.com/gin-gonic/gin"
)

// --- get ---

type GetRequestFields struct {
	CanvasID uint `json:"canvas_id" check:"not null; more: 0"`
}

func (c *GetRequestFields) JSON(ctx *gin.Context) error {
	return ctx.BindJSON(c)
}

type GetResponseFields struct {
	Header ginUtils.BaseRespHeader `json:"header"`
	Cover  string                  `json:"cover"`
}

// --- update ---

type UpdateRequestFields struct {
	CanvasID uint   `json:"canvas_id" check:"not null; more: 0"`
	Data     string `json:"data"`
}

func (c *UpdateRequestFields) JSON(ctx *gin.Context) error {
	return ctx.BindJSON(c)
}

type UpdateResponseFields struct {
	Header ginUtils.BaseRespHeader `json:"header"`
}
