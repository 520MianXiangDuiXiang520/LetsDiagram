package canvas

import (
	"encoding/json"
	"github.com/520MianXiangDuiXiang520/ginUtils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"lets_diagram/src"
	"lets_diagram/src/core"
	"lets_diagram/src/core/access"
	"lets_diagram/src/core/connection"
	"lets_diagram/src/core/connection/handler"
	"lets_diagram/src/core/connection/pingpong"
	"lets_diagram/src/core/cooperate"
	"lets_diagram/src/core/versionlink"
	"lets_diagram/src/dao"
	"lets_diagram/src/models"
	"lets_diagram/src/models/nottable"
	"log"
	"net/http"
	"strconv"
)

func checkUser(request *ginUtils.Request, response *ginUtils.Response) (*models.User, bool) {
	user, ok := request.Ctx.Get("user")
	if !ok {
		response.RespCode = http.StatusUnauthorized
		response.Resp = NewCanvasResponseField{
			Header:   ginUtils.UnauthorizedRespHeader,
			CanvasID: 0,
		}
	}
	return user.(*models.User), ok
}

func RenameLogic(request *ginUtils.Request, response *ginUtils.Response) error {
	req := request.Req.(*RenameRequestFields)
	if !dao.CanvasRename(req.Name, req.CanvasID) {
		response.RespCode = http.StatusBadRequest
		response.Resp = RenameResponseFields{Header: ginUtils.ParamErrorRespHeader}
		return nil
	}
	response.RespCode = http.StatusOK
	response.Resp = RenameResponseFields{Header: ginUtils.SuccessRespHeader}
	return nil
}

// NewLogic 处理用户新建 Canvas 请求
func NewLogic(request *ginUtils.Request, response *ginUtils.Response) error {
	user, ok := checkUser(request, response)
	req := request.Req.(*NewCanvasRequestField)
	if !ok {
		return nil
	}
	name := req.Name
	if name == "" {
		name = "未命名文件"
	}
	canvasInfo, ok := dao.NewCanvas(name, user.ID)
	if !ok {
		response.RespCode = http.StatusInternalServerError
		response.Resp = NewCanvasResponseField{
			Header:   ginUtils.SystemErrorRespHeader,
			CanvasID: 0,
		}
		return nil
	}
	response.Resp = NewCanvasResponseField{
		Header:   ginUtils.SuccessRespHeader,
		CanvasID: canvasInfo.ID,
	}
	return nil
}

// OpenLogic 从返回已经持久化了的 Canvas 数据，同时初始化版本链并且将执行此操作的用户设置
// 为该 canvas 的 root 用户。
func OpenLogic(request *ginUtils.Request, response *ginUtils.Response) error {
	req := request.Req.(*OpenCanvasRequestField)
	user, ok := checkUser(request, response)
	if !ok {
		return nil
	}
	canvas, err := dao.GetCanvasData(req.CanvasID)
	if err != nil {
		response.RespCode = http.StatusBadRequest
		response.Resp = OpenCanvasResponseField{Header: ginUtils.ParamErrorRespHeader}
		return nil
	}
	versionlink.GetVersionsLink().NewVersionsLink(req.CanvasID, canvas.Version)
	// 初始化访问控制信息, 授予 root 权限
	if !stwOpen(req.CanvasID) {
		access.GetAccessControlTable().InitPermission(req.CanvasID, user.ID)
	}
	response.RespCode = http.StatusOK
	response.Resp = OpenCanvasResponseField{
		Header:    ginUtils.SuccessRespHeader,
		Data:      canvas.Data,
		Version:   canvas.Version,
		Authority: access.CanvasJurisdictionMarkRoot,
	}
	return nil
}

func stwOpen(canvasID uint) bool {
	if !access.GetAccessControlTable().IsSTW(canvasID) {
		return false
	}
	g, ok := connection.GetWSPool().Load(canvasID)
	if !ok {
		return false
	}
	g.RootRejoinSign <- struct{}{}
	return true
}

func JoinLogic(request *ginUtils.Request, response *ginUtils.Response) error {
	req := request.Req.(*JoinRequestFields)
	user, ok := checkUser(request, response)
	if !ok {
		return nil
	}
	canvasID, ok := cooperate.GetCooperateCodeTable().GetCanvasID(req.Code)
	if !ok {
		response.RespCode = http.StatusBadRequest
		response.Resp = JoinResponseFields{
			Header: ginUtils.BaseRespHeader{
				Code: http.StatusBadRequest,
				Msg:  "请输入正确的协作码！",
			},
		}
		return nil
	}
	canvas, err := dao.GetCanvasData(canvasID)
	if err != nil {
		response.RespCode = http.StatusBadRequest
		response.Resp = JoinResponseFields{
			Header: ginUtils.ParamErrorRespHeader,
		}
		return nil
	}
	access.GetAccessControlTable().SetPermission(access.RootSetPermission,
		canvasID, user.ID, access.CanvasJurisdictionMarkReadOnly)
	connection.SendAllPong(canvasID, pingpong.GetSimpleNotifyPong(pingpong.PongSimpleNotifyUserAdd, user))
	response.RespCode = http.StatusOK
	response.Resp = JoinResponseFields{
		Header:    ginUtils.SuccessRespHeader,
		Data:      canvas.Data,
		Version:   canvas.Version,
		CanvasID:  canvasID,
		Authority: access.CanvasJurisdictionMarkReadOnly,
	}
	return nil
}

// AllLogic 返回请求者持有的所有 canvas
func AllLogic(request *ginUtils.Request, response *ginUtils.Response) error {
	user, ok := checkUser(request, response)
	if !ok {
		return nil
	}
	req := request.Req.(*AllCanvasRequestField)
	page, size := req.Page, req.Size
	if size <= 0 || page <= 0 {
		size = src.GetSetting().DefaultSetting.DefaultPageSize
		page = src.GetSetting().DefaultSetting.DefaultPage
	}
	total := dao.GetUserAllCanvasTotal(user.ID)
	if total == 0 {
		response.RespCode = http.StatusOK
		response.Resp = AllCanvasResponseField{
			Header: ginUtils.SuccessRespHeader,
			List:   []nottable.FullCanvas{},
			Total:  0,
		}
		return nil
	}
	canvas := dao.GetUserAllCanvas(user.ID, (page-1)*size, size)
	if canvas == nil {
		response.RespCode = http.StatusBadRequest
		response.Resp = AllCanvasResponseField{
			Header: ginUtils.ParamErrorRespHeader,
		}
		return nil
	}

	response.RespCode = http.StatusOK
	response.Resp = AllCanvasResponseField{
		Header: ginUtils.SuccessRespHeader,
		List:   canvas,
		Total:  total,
	}
	return nil
}

// CooperateLogic 处理 root 用后开启协作时的请求，主要工作是为每个 canvas 生成一个唯一的协作码
func CooperateLogic(request *ginUtils.Request, response *ginUtils.Response) error {
	req := request.Req.(*CooperateRequestField)
	code := cooperate.GetCooperateCodeTable().GetCode(req.CanvasID)
	response.RespCode = http.StatusOK
	response.Resp = CooperateResponse{
		Header:        ginUtils.SuccessRespHeader,
		CooperateCode: code,
	}
	return nil
}

func CheckCooperateLogic(request *ginUtils.Request, response *ginUtils.Response) error {
	req := request.Req.(*CheckCooperateRequestFields)
	_, ok := cooperate.GetCooperateCodeTable().GetCanvasID(req.Code)
	if !ok {
		response.RespCode = http.StatusBadRequest
		response.Resp = CheckCooperateResponseFields{
			Header: ginUtils.ParamErrorRespHeader,
			Result: false,
		}
		return nil
	}
	response.RespCode = http.StatusOK
	response.Resp = CheckCooperateResponseFields{
		Header: ginUtils.SuccessRespHeader,
		Result: true,
	}
	return nil
}

func getSimpleUserList(ids []uint) []nottable.SimpleUser {
	result := make([]nottable.SimpleUser, len(ids))
	for i, id := range ids {
		user, ok := dao.GetUserByID(id)
		if ok {
			result[i] = nottable.SimpleUser{
				ID:       user.ID,
				Username: user.Username,
				Email:    user.Email,
			}
		}
	}
	return result
}

// CollaboratorsLogic 响应当前 canvas 的所有协作者
func CollaboratorsLogic(request *ginUtils.Request, response *ginUtils.Response) error {
	req := request.Req.(*CollaboratorsRequestFields)
	canvasID := req.CanvasID
	listDict := access.GetAccessControlTable().GetList(canvasID)
	if listDict == nil {
		response.RespCode = http.StatusBadRequest
		response.Resp = CollaboratorsResponseFields{
			Header: ginUtils.ParamErrorRespHeader,
		}
		return nil
	}
	rooters := getSimpleUserList(listDict[access.CanvasJurisdictionMarkRoot])
	writers := getSimpleUserList(listDict[access.CanvasJurisdictionMarkWrite])
	managers := getSimpleUserList(listDict[access.CanvasJurisdictionMarkManager])
	readers := getSimpleUserList(listDict[access.CanvasJurisdictionMarkReadOnly])
	response.RespCode = http.StatusOK
	response.Resp = CollaboratorsResponseFields{
		Header: ginUtils.SuccessRespHeader,
		Rooter: rooters,
		Manger: managers,
		Writer: writers,
		Reader: readers,
	}
	return nil
}

func PaintLogic(ws *websocket.Conn, ctx *gin.Context) {
	u, ok := ctx.Get("user")
	if !ok {
		log.Printf("[websocket] can not get user: request = %v", ctx.Request)
		return
	}
	user := u.(*models.User)
	canvasID := ctx.Query("canvas_id")
	cID, err := strconv.Atoi(canvasID)
	if err != nil {
		log.Printf("[websocket] can not get canvasID: request = %v", ctx.Request)
		return
	}
	log.Printf("the user(%s:%d) opened the canvas(%d)", user.Username, user.ID, cID)
	defer func() {
		log.Printf("the user(%s:%d) leave the canvas(%d)", user.Username, user.ID, cID)
		core.ConnectionClose(uint(cID), user)
	}()

	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		switch messageType {
		case websocket.CloseMessage:
			break
		case websocket.TextMessage:
			var ping pingpong.PingMessage
			err := json.Unmarshal(message, &ping)
			if err != nil {
				log.Printf("Unmarshal fail: %v", err)
			}
			switch ping.Type {
			case pingpong.PingTypeVersionControl:
				handler.VersionControlHandler(message, uint(cID), user, ws)
			case pingpong.PingTypeHeartbeat:
				handler.HeartBeatHandler(message, uint(cID), user, ws)
			case pingpong.PingTypePermissionControl:
				handler.PermissionHandler(message, uint(cID), user, ws)
			}
		}
	}
}

func ForkLogic(request *ginUtils.Request, response *ginUtils.Response) error {
	return nil
}

func DeleteLogic(request *ginUtils.Request, response *ginUtils.Response) error {
	req := request.Req.(*DeleteRequestFields)
	user, _ := checkUser(request, response)
	// 检查是否有未持久化的版本
	vl, ok := versionlink.GetVersionsLink().Load(req.CanvasID)
	if ok {
		vl.EnduranceMutex.Lock()
		vl.SetSkipEndurance()
		nodes, count := vl.GetLastN(vl.Length)
		versionlink.NewEndurance(nodes, count, req.CanvasID)
		vl.EnduranceMutex.Unlock()
		versionlink.GetVersionsLink().Delete(req.CanvasID)
	}
	if !dao.DeleteCanvas(req.CanvasID, user.ID) {
		response.RespCode = http.StatusBadRequest
		response.Resp = DeleteResponseFields{Header: ginUtils.ParamErrorRespHeader}
		return nil
	}
	response.RespCode = http.StatusOK
	response.Resp = DeleteResponseFields{Header: ginUtils.SuccessRespHeader}
	return nil
}
