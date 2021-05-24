package dao

import (
    "github.com/520MianXiangDuiXiang520/GoTools/dao"
    `lets_diagram/src`
    "lets_diagram/src/dao/mysql"
    "lets_diagram/src/models"
    "log"
    "time"
)

func GetUserByToken(token string) (*models.User, bool) {
	// TODO: use redis
	return mysql.SelectUserByToken(token)
}

func GetUserByMailPsw(email, password string) (*models.User, bool) {
	return mysql.SelectUserByEmailPSW(email, password)
}

func GetUserByID(id uint) (*models.User, bool) {
	return mysql.SelectUserByID(id)
}

func InsertUserToken(userID uint, token string) bool {
	// TODO: use redis
	return mysql.InsertUserToken(userID, token)
}

func UpdateUserTokenExpireTime(token string) bool {
	// TODO: update redis
	return mysql.UpdateUserTokenExpireTime(token,
		time.Now().UnixNano()+src.GetSetting().TokenSetting.ExpireTime)
}

// IsEmailDuplicate if the email has already been registered, return true
func IsEmailDuplicate(email string) bool {
	db := dao.GetDB()
	return mysql.IsEmailDuplicate(db, email)
}

func InsertNewUser(username, password, email string) bool {
	db := dao.GetDB()
	newUser := &models.User{
		Username: username,
		Password: password,
		Email:    email,
	}
	err := mysql.InsertNewUser(db, newUser)
	if err != nil {
		log.Printf("dao/user/insterNewUser %v \n", err)
		return false
	}
	return true
}
