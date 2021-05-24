package main

import (
    "github.com/520MianXiangDuiXiang520/GoTools/dao"
    `lets_diagram/src`
    `lets_diagram/src/models`
)

func init() {
	src.InitSetting("../../../config/setting.json")
	dao.InitDBSetting(src.GetSetting().MySQLSetting)
}

func createTable() {
	// dao.GetDB().CreateTable(&models.User{})
	// dao.GetDB().CreateTable(&models.UserToken{})
	// dao.GetDB().CreateTable(&models.Canvas{})
	// dao.GetDB().CreateTable(&models.UserCanvas{})
	// dao.GetDB().CreateTable(&models.Covers{})
    dao.GetDB().AutoMigrate(&models.CanvasData{})
}

func main() {
	createTable()
}
