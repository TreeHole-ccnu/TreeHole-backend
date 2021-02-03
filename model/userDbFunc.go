package model

import (
	//"github.com/TreeHole-ccnu/TreeHole-backend/util"

	"log"
)

//返回0，该用户已注册，返回1，该用户未注册
func ConfirmPhone(phone string) int {
	var user User
	if err := Db.Self.Model(&User{}).Where(&User{Phone: phone}).First(&user).Error; err != nil {
		return 1
	}
	return 0
}

//i = 2时，该用户输入密码错误、i = 3时，该用户输入密码正确
func ConfirmUserPhone(phone string, password string) (int, error, int) {
	var u User

	if err := Db.Self.Model(&User{}).Where(&User{Phone: phone}).First(&u).Error; err != nil {
		return 0, err, -1
	}
	if password != u.Password {
		return 2, nil, u.Level
	} else {
		return 3, nil, u.Level
	}
}

//查看验证码是否正确，1--正确，0--不正确
func ConfirmUserVcd(phone string, code string) int {
	var flag int
	getCode := GetRedis(phone)

	if code == getCode {
		flag = 1
	} else {
		flag = 0
	}
	return flag
}

func CreateUserRegisterInfo(phone string, password string) error {
	user := &User{
		Phone:          phone,
		Password:       password,
		Id:             0,
		Level:          0,
		Name:           "",
		Sex:            "",
		Birth:          "",
		Nation:         "",
		NativePlace:    "",
		Email:          "",
		IdentityNumber: "",
		ImageUrl:       "",
	}

	if err := Db.Self.Model(&User{}).Create(&user).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}
