package model

import (
	//"github.com/TreeHole-ccnu/TreeHole-backend/util"

	//"log"
)

func ResetPassword (phone string,password string )error{
	if  err :=Db.Self.Model(&User{}).Where(&User{Phone:phone}).Update("password", &password).Error; err != nil {
		return err
	}
	return nil
}

func ResetNormalInfo(user User)error{
	
	if err := Db.Self.Model(&User{}).Where(&User{Phone:user.Phone}).Update(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetInfo (phone string)(l User,err error){
	if  err :=Db.Self.Model(&User{}).Where(&User{Phone:phone}).First(&l).Error; err != nil {
		return l,err
	}
	return l,nil
}

func Image_modify(phone string, image_url string)error {
	if err :=Db.Self.Model(&User{}).Where(&User{Phone:phone}).Update("image_url", &image_url).Error; err != nil {
		return err
	}
	return nil
}