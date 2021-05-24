package middleware

import (
	"bytes"
	"github.com/520MianXiangDuiXiang520/ginUtils/middleware"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"lets_diagram/src/dao/mysql"
	"log"
)

// ISCreator 检查当前用户是不是 Canvas 的创建者
func ISCreator(context *gin.Context) bool {
	b, err := ioutil.ReadAll(context.Request.Body)
	defer func() {
		context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	}()
	context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	type req struct {
		CanvasID uint `json:"canvas_id"`
	}
	r := &req{}
	u, ok := context.Get("user")
	if !ok {
		return false
	}
	user := u.(middleware.UserBase)
	err = context.BindJSON(&r)
	if err != nil {
		log.Println("bind error", err)
		return false
	}
	_, err = mysql.SelectByUserIDCanvasID(r.CanvasID, uint(user.GetID()))
	return err == nil
}
