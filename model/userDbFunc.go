package model

import (
	//"github.com/TreeHole-ccnu/TreeHole-backend/util"

	"log"
)

//i = 1时，该用户还未注册、i = 2时，该用户输入密码错误、i = 3时，该用户输入密码正确
func ConfirmUserPhone(phone string, password string) (int, error) {
	var realPassword string

	if Db.Self.Model(&User{}).Where(&User{Phone:phone}).RecordNotFound() {
		return 1, nil
	}

	if err := Db.Self.Model(&User{}).Where(&User{Phone:phone}).Pluck("password", &realPassword).Error; err != nil {
		return 0, err
	}
	if password != realPassword {
		return 2, nil
	} else {
		return 3, nil
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

func CreateUserRegisterInfo (phone string, password string) error{
	user := &User{
		Phone:          phone,
		Passoword:      password,
		Id:             0,
		Level:          0,
		Name:           "",
		Sex:            "",
		Birth:          "",
		Native:         "",
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