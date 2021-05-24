package models

import (
	"github.com/jinzhu/gorm"
)

// 用户表
type User struct {
	gorm.Model
	Username string `gorm:"size:16;not null;unique"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null"`
}

type UserToken struct {
	gorm.Model
	UserID     uint   `gorm:"not null"`
	Token      string `gorm:"not null"`
	ExpireTime int64  `gorm:"not null"`
}

func (u User) GetID() int {
	return int(u.ID)
}
