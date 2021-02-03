package model

import (
//"github.com/TreeHole-ccnu/TreeHole-backend/util"

//"log"
)

//返回0，该用户不存在（还未注册）；返回1，该用户存在
func ConfirmUser(phone string) int {
	if Db.Self.Model(&User{}).Where(&User{Phone: phone}).RecordNotFound() {
		return 0
	}
	return 1
}

func ResetPassword(phone string, password string) error {
	if err := Db.Self.Model(&User{}).Where(&User{Phone: phone}).Update("password", password).Error; err != nil {
		return err
	}
	return nil
}

func ResetNormalInfo(user User) error {
	if err := Db.Self.Model(&User{}).Where(&User{Phone: user.Phone}).Update(User{
		Phone:          user.Phone,
		Name:           user.Name,
		Sex:            user.Sex,
		Birth:          user.Birth,
		Nation:         user.Nation,
		NativePlace:    user.NativePlace,
		Email:          user.Email,
		IdentityNumber: user.IdentityNumber,
		ImageUrl:       user.ImageUrl,
	}).Error; err != nil {
		return err
	}
	return nil
}

func GetInfo(phone string) (l User, err error) {
	if err := Db.Self.Model(&User{}).Where(&User{Phone: phone}).First(&l).Error; err != nil {
		return l, err
	}
	return l, nil
}

func Image_modify(phone string, image_url string) error {
	if err := Db.Self.Model(&User{}).Where(&User{Phone: phone}).Update("image_url", image_url).Error; err != nil {
		return err
	}
	return nil
}
