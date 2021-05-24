package models

import (
	"github.com/jinzhu/gorm"
)

type CanvasData struct {
	gorm.Model
	Data     string `json:"data" gorm:"column:data"`
	Version  int64  `json:"version" gorm:"column:version"`  // 最后一次持久化的版本号
}

type Canvas struct {
	gorm.Model
	Name      string // 画布名
	CoverID   uint   `json:"cover_id" gorm:"column:cover_id"`// 封面图的路径
	DataID    uint   // 对应 CanvasData 表的 ID
}

// UserCanvas 用来存储某个用户有哪些自己的 canvas
type UserCanvas struct {
	gorm.Model
	UserID   uint `gorm:"not null"`
	CanvasID uint `gorm:"not null"`
}

type Covers struct {
    gorm.Model
    Data     string `json:"data" gorm:"column:data"`
}