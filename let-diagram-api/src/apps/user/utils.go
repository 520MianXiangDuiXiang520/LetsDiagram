package user

import (
	"github.com/520MianXiangDuiXiang520/GoTools/crypto"
	"time"
)

func GetToken(email, psw string) string {
	return crypto.MD5([]string{
		email, psw, time.Now().String(),
	})
}
