package user

import (
	"github.com/520MianXiangDuiXiang520/GoTools/crypto"
	"github.com/520MianXiangDuiXiang520/ginUtils"
	"lets_diagram/src/dao"
	"net/http"
)

func LoginLogic(request *ginUtils.Request, response *ginUtils.Response) error {
	loginRequest := request.Req.(*LoginRequestFields)
	psw := crypto.SHA256([]string{loginRequest.Password})
	user, ok := dao.GetUserByMailPsw(loginRequest.Email, psw)
	if !ok {
		response.Resp = LoginResponseField{
			Header: ginUtils.BaseRespHeader{
				Code: http.StatusBadRequest,
				Msg:  "用户名或密码错误！",
			},
			Token: "",
		}
		return nil
	}

	token := GetToken(loginRequest.Email, psw)
	ok = dao.InsertUserToken(user.ID, token)
	if !ok {
		response.RespCode = http.StatusInternalServerError
		response.Resp = LoginResponseField{
			Header: ginUtils.SystemErrorRespHeader,
			Token:  "",
		}
		return nil
	}

	response.Resp = LoginResponseField{
		Header: ginUtils.SuccessRespHeader,
		Token:  token,
	}
	return nil
}

func RegisterLogic(request *ginUtils.Request, response *ginUtils.Response) error {
	registerRequest := request.Req.(*RegisterRequestFields)
	if dao.IsEmailDuplicate(registerRequest.Email) {
		response.RespCode = http.StatusBadRequest
		response.Resp = RegisterResponseFields{Header: ginUtils.BaseRespHeader{
			Code: http.StatusBadRequest,
			Msg:  "您的邮箱已经注册过了，直接去登陆吧！",
		}}
		return nil
	}
	psw := crypto.SHA256([]string{registerRequest.Password})
	ok := dao.InsertNewUser(registerRequest.Username, psw, registerRequest.Email)
	if !ok {
		response.RespCode = http.StatusInternalServerError
		response.Resp = RegisterResponseFields{Header: ginUtils.SystemErrorRespHeader}
		return nil
	}
	response.RespCode = http.StatusOK
	response.Resp = RegisterResponseFields{
		Header: ginUtils.SuccessRespHeader,
	}
	return nil
}
