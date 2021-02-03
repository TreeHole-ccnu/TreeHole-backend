package model

//通过手机号查找志愿者
func SearchVolunteer(phone string) (string, int, error) {
	var v Volunteer
	if err := Db.Self.Model(&Volunteer{}).Where(&Volunteer{Phone: phone}).First(&v).Error; err != nil {
		return "", 0, err
	}
	return v.Name, v.Id, nil
}

//获取志愿者详细信息
func GetVolunteerInfo(id int) (Volunteer, error) {
	var v Volunteer
	if err := Db.Self.Model(&Volunteer{}).Where(&Volunteer{Id: id}).Find(&v).Error; err != nil {
		return v, err
	}
	return v, nil
}

//获取支援者简历详细信息
func GetUserResume(id int) ([]Resume, error) {
	var r []Resume
	if err := Db.Self.Model(&Resume{}).Where(&Resume{Id: id}).Find(&r).Error; err != nil {
		return r, err
	}
	return r, nil
}
