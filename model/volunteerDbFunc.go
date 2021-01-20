package model

import "log"

//创建志愿者申请表单
func CreateVolunteer(l Volunteer) error {
	if err := Db.Self.Model(&Volunteer{}).Create(&l).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//获取志愿者表申请进度
//1--完善个人信息 2--提交申请表 3--审核申报人信息 4--志愿者协会审核 5--审核通过
func GetCheckId(phone string) (int, error) {
	var i int
	if err := Db.Self.Model(&Volunteer{}).Where(&Volunteer{Phone: phone}).Pluck("ischeck", &i).Error; err != nil {
		return 0, err
	}
	return i, nil
}

//返回0，该用户还未申请志愿者；返回1，该用户已申请志愿者
func ConfirmVolunteer(phone string) int {
	if Db.Self.Model(&Volunteer{}).Where(&Volunteer{Phone: phone}).RecordNotFound() {
		return 0
	}
	return 1
}
