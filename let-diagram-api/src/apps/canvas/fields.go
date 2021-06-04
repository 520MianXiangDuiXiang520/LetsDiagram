package canvas

import (
	"github.com/520MianXiangDuiXiang520/ginUtils"
	"github.com/gin-gonic/gin"
	"lets_diagram/src/core/access"
	"lets_diagram/src/models/nottable"
)

type RenameRequestFields struct {
	CanvasID uint   `json:"canvas_id" check:"not null; more: 0"`
	Name     string `json:"name" check:"len: [0, 128]"`
}

func (c *RenameRequestFields) JSON(ctx *gin.Context) error {
	return ctx.BindJSON(c)
}

type RenameResponseFields struct {
	Header ginUtils.BaseRespHeader `json:"header"`
}

// --- delete ---

type DeleteRequestFields struct {
	CanvasID uint `json:"canvas_id" check:"not null; more: 0"`
}

func (c *DeleteRequestFields) JSON(ctx *gin.Context) error {
	return ctx.BindJSON(c)
}

type DeleteResponseFields struct {
	Header ginUtils.BaseRespHeader `json:"header"`
}

// --- fork ---

type ForkRequestFields struct {
	BaseCanvasID uint `json:"base_canvas_id" check:"not null; more: 0"`
}

func (f *ForkRequestFields) JSON(ctx *gin.Context) error {
	return ctx.BindJSON(f)
}

type ForkResponseFields struct {
	Header   ginUtils.BaseRespHeader `json:"header"`
	CanvasID uint                    `json:"canvas_id"`
}

// --- collaborators ---

type CollaboratorsRequestFields struct {
	CanvasID uint `json:"canvas_id" check:"not null"`
}

func (f *CollaboratorsRequestFields) JSON(ctx *gin.Context) error {
	return ctx.BindJSON(f)
}

type CollaboratorsResponseFields struct {
	Header ginUtils.BaseRespHeader `json:"header"`
	Rooter []nottable.SimpleUser   `json:"rooter"`
	Manger []nottable.SimpleUser   `json:"manger"`
	Writer []nottable.SimpleUser   `json:"writer"`
	Reader []nottable.SimpleUser   `json:"reader"`
}

// --- checkCooperate ---

type CheckCooperateRequestFields struct {
	Code string `json:"code" check:"not null"`
}

func (f *CheckCooperateRequestFields) JSON(ctx *gin.Context) error {
	return ctx.BindJSON(f)
}

type CheckCooperateResponseFields struct {
	Header ginUtils.BaseRespHeader `json:"header"`
	Result bool                    `json:"result"`
}

// --- join ---

type JoinRequestFields struct {
	Code string `json:"code" check:"not null"` // 协作码
}

func (f *JoinRequestFields) JSON(ctx *gin.Context) error {
	return ctx.BindJSON(f)
}

type JoinResponseFields struct {
	Header    ginUtils.BaseRespHeader       `json:"header"`
	CanvasID  uint                          `json:"canvas_id"` // 协作的 canvasID
	Data      interface{}                   `json:"data"`      // 同步持久化数据
	Version   int64                         `json:"version"`   // 用来同步未持久化的数据
	Authority access.CanvasJurisdictionMark `json:"authority"` // 初始权限，ReadOnly
}

// --- cooperate ---

type CooperateRequestField struct {
	CanvasID uint `json:"canvas_id" check:"not null"`
}

func (f *CooperateRequestField) JSON(ctx *gin.Context) error {
	return ctx.BindJSON(f)
}

type CooperateResponse struct {
	Header        ginUtils.BaseRespHeader `json:"header"`
	CooperateCode string                  `json:"cooperate_code"`
}

// --- new ---

type NewCanvasRequestField struct {
	Name string `json:"name" check:"len: [0, 128]"`
}

func (f *NewCanvasRequestField) JSON(ctx *gin.Context) error {
	return ctx.BindJSON(f)
}

type NewCanvasResponseField struct {
	Header   ginUtils.BaseRespHeader `json:"header"`
	CanvasID uint                    `json:"canvas_id"`
}

// --- open ---

type OpenCanvasRequestField struct {
	CanvasID uint `json:"canvas_id" check:"not null; more: 0"`
}

func (f *OpenCanvasRequestField) JSON(ctx *gin.Context) error {
	return ctx.BindJSON(f)
}

type OpenCanvasResponseField struct {
	Header    ginUtils.BaseRespHeader       `json:"header"`
	Data      interface{}                   `json:"data"`
	Version   int64                         `json:"version"`
	Authority access.CanvasJurisdictionMark `json:"authority"`
}

// --- all ---

type AllCanvasRequestField struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

func (f *AllCanvasRequestField) JSON(ctx *gin.Context) error {
	return ctx.BindJSON(f)
}

type AllCanvasRespList struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type AllCanvasResponseField struct {
	Header ginUtils.BaseRespHeader `json:"header"`
	Total  int                     `json:"total"`
	List   []nottable.FullCanvas   `json:"list"`
}

// --- paint ---

type PaintRequestField struct {
	CanvasID uint   `json:"canvas_id" check:"not null"`
	Token    string `json:"token"`
}

func (f *PaintRequestField) JSON(ctx *gin.Context) error {
	return ctx.BindJSON(f)
}

type PaintResponseField struct {
}

// PaintWSRequestFields 表示 websocket 请求时需要携带的参数
type PaintWSRequestFields struct {
	ClientVersion int64  `json:"version"` // 客户端所处的版本
	Token         string `json:"token"`
	CanvasID      uint   `json:"canvas_id"` // 要操作的 canvas
}
