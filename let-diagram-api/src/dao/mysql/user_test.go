package mysql

import (
    "github.com/520MianXiangDuiXiang520/GoTools/dao"
    `lets_diagram/src`
    "lets_diagram/src/models"
    "strconv"
    "testing"
    "time"
)

func init() {
	src.InitSetting("../../../config/setting.json")
	_ = dao.InitDBSetting(src.GetSetting().MySQLSetting)
}

func RandomEmail() string {
	start := strconv.FormatInt(time.Now().UnixNano(), 10)
	return start + "@junebao.top"
}

func TestIsEmailDuplicate(t *testing.T) {
	db := dao.GetDB()
	testUser := &models.User{
		Username: "test01",
		Password: "0x111",
		Email:    RandomEmail(),
	}
	defer func() {
		db.Delete(&testUser)
	}()
	if ret := InsertNewUser(db, testUser); ret != nil {
		t.Error("fail to insert new user")
	}

	if ret := IsEmailDuplicate(db, testUser.Email); !ret {
		t.Error("fail test")
	}
}
