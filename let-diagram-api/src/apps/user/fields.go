package user

import (
	"github.com/520MianXiangDuiXiang520/ginUtils"
	"github.com/gin-gonic/gin"
)

type LoginRequestFields struct {
	Email    string `json:"email" check:"email"`
	Password string `json:"password" check:"len: [7, 20]"`
}

func (f *LoginRequestFields) JSON(ctx *gin.Context) error {
	return ctx.BindJSON(f)
}

type LoginResponseField struct {
	Header ginUtils.BaseRespHeader `json:"header"`
	Token  string                  `json:"token"`
}

type RegisterRequestFields struct {
	Username string `json:"username" check:"len: [5, 16]; not null"`
	Email    string `json:"email" check:"email; not null"`
	Password string `json:"password" check:"len: [7, 20]; not null"`
}

func (f *RegisterRequestFields) JSON(ctx *gin.Context) error {
	return ctx.BindJSON(f)
}

type RegisterResponseFields struct {
	Header ginUtils.BaseRespHeader `json:"header"`
}
