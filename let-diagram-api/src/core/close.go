package core

import (
	"lets_diagram/src/core/access"
	"lets_diagram/src/core/connection"
	"lets_diagram/src/core/connection/pingpong"
	"lets_diagram/src/core/cooperate"
	"lets_diagram/src/models"
	"log"
	"time"
)

// DestroyCanvas 销毁与 canvas 相关的所有内存资源
// 包括连接信息，协作码，协作者权限控制表
func DestroyCanvas(canvasID uint) {
	connection.GetWSPool().RemoveCanvas(canvasID)
	cooperate.GetCooperateCodeTable().Remove(canvasID)
	access.GetAccessControlTable().RemoveCanvas(canvasID)
	log.Printf("canvas %d 已销毁", canvasID)
}

func RootClose(canvasID, userID uint) {
	access.GetAccessControlTable().STW(canvasID)
	_ = connection.GetWSPool().RemoveConnect(canvasID, userID)
	connection.STWPongNotifyAll(canvasID, pingpong.PongRootCloseSTWCrash)
	g, ok := connection.GetWSPool().Load(canvasID)
	if ok {
		g.RootCloseWaiter = time.NewTimer(time.Minute)
		select {
		case <-g.RootCloseWaiter.C:
			DestroyCanvas(canvasID)
			log.Printf("canvas(%d) is destroyed", canvasID)
		case <-g.RootRejoinSign:
			log.Printf("the user(%d) rejoin the canvas(%d)", userID, canvasID)
			access.GetAccessControlTable().UnSTW(canvasID)
			g.RootCloseWaiter.Stop()
			connection.STWPongNotifyAll(canvasID, pingpong.PongRootCloseSTWRecovery)
			return
		}
	}
}

func OrdinaryClose(canvasID uint, user *models.User) {
	_ = connection.GetWSPool().RemoveConnect(canvasID, user.ID)
	access.GetAccessControlTable().DriveOut(canvasID, user.ID)
	connection.SendAllPong(canvasID, pingpong.GetSimpleNotifyPong(pingpong.PongSimpleNotifyUserOut, user))
	log.Printf("canvas %d ~ user %d 已退出", canvasID, user.ID)
}

// ConnectionClose 执行 socket 关闭后的善后操作
// 1. 任何情况下，客户端关闭连接（包括刷新），服务端都会删除该连接
// 2. 画布创建者关闭连接后，等待一分钟，如果创建者未重新加入进来，
//    销毁与该画布相关的所有连接，包括与之相关的权限控制表，协作码
//    等信息，且在等待期间，剥夺所有用户的写权限，直到创建者重新加入
//    或超时后退出。
// 3. 普通协作者（通过协作码加入的用户）关闭连接后，直接删除与该用户
//    相关的连接，权限信息
func ConnectionClose(canvasID uint, user *models.User) {
	root, ok := access.GetAccessControlTable().GetRoot(canvasID)
	if ok && root == user.ID {
		RootClose(canvasID, user.ID)
	} else {
		OrdinaryClose(canvasID, user)
	}
}
