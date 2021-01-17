package model

import (
	//"github.com/TreeHole-ccnu/TreeHole-backend/util"

	"log"
)

func ResetPassword (phone string,password string )error{
	if  err :=Db.Self.Model(&User{}).Where(&User{Phone:phone}).Update("password", &password).Error; err != nil {
		return err
	}
	return nil
}

func CreateUserNormalInfo(user User)error{
	if err := Db.Self.Model(&User{}).Where(&User{Phone:user.Phone}).Update(&user).Error; err != nil {
		return err
	}
	return nil
}