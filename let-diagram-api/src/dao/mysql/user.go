package mysql

import (
    "github.com/520MianXiangDuiXiang520/GoTools/dao"
    "github.com/jinzhu/gorm"
    `lets_diagram/src`
    "lets_diagram/src/models"
    "log"
    "time"
)

func SelectUserByID(userID uint) (*models.User, bool) {
	user := models.User{}
	db := dao.GetDB().Where("id = ?", userID).First(&user)
	if db.Error != nil {
		log.Printf("[dao/mysql] fail to get user info by userID(%d), \nerror info: %v \n", userID, db.Error)
		return nil, false
	}
	return &user, true
}

func SelectUserTokenByToken(token string) (*models.UserToken, bool) {
	userToken := models.UserToken{}
	db := dao.GetDB().Where("token = ?", token).First(&userToken)
	if db.Error != nil {
		log.Printf("[dao/mysql] fail to get userToken by token, err info: \n %v \n", db.Error)
		return nil, false
	}
	return &userToken, true
}

func SelectUserByToken(token string) (*models.User, bool) {
	userToken, ok := SelectUserTokenByToken(token)
	// if the token has expired
	if !ok || userToken.ExpireTime < time.Now().UnixNano() {
		return nil, false
	}
	return SelectUserByID(userToken.UserID)
}

func SelectUserByEmailPSW(email, password string) (*models.User, bool) {
	user := models.User{}
	db := dao.GetDB().Where("email = ? AND password = ?", email, password).First(&user)
	if db.Error != nil {
		log.Printf("[dao/mysql] fail to get user info by email(%s) and password(%s), erroe info: \n %v", email, password, db.Error)
		return nil, false
	}
	return &user, true
}

func InsertUserToken(userID uint, token string) bool {
	db := dao.GetDB().Create(&models.UserToken{
		UserID:     userID,
		Token:      token,
		ExpireTime: time.Now().UnixNano() + src.GetSetting().TokenSetting.ExpireTime,
	})
	if db.Error != nil {
		log.Printf("[dao/mysql] fail to insert userToken({userID: %d, Token: %s}), error info: \n %v", userID, token, db.Error)
		return false
	}
	return true
}

func UpdateUserTokenExpireTime(token string, expireTime int64) bool {
	d := dao.GetDB().Model(&models.UserToken{}).Where("token = ?",
		token).Update("expire_time", expireTime)
	if d.Error != nil {
		log.Printf("[dao/mysql/user] Failed to update the expire time of the row with token = %s to %d, error info: \n %v \n", token, expireTime, d.Error)
		return false
	}
	return true
}

func InsertNewUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}

func IsEmailDuplicate(db *gorm.DB, email string) bool {
	e := &models.User{}
	err := db.Select("email").Where("email = ?", email).Find(e).Error
	if err != nil {
		return false
	}
	return e.Email == email
}

func SelectUserByEmail(db *gorm.DB, email string) (*models.User, bool) {
	user := models.User{}
	d := db.Where("email = ?", email).Find(&user)
	if d.Error != nil {
		log.Printf("[dao/mysql/user] fail to get users by email(%s), error info: \n %v \n", email, d.Error)
		return nil, false
	}
	return &user, true
}
