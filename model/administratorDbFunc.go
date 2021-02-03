package model

import "fmt"

//返回0，该用户不存在（还未注册）；返回1，该用户存在
func CheckUser(id int) bool {
	if Db.Self.Model(&User{}).Where(&User{Id: id}).RecordNotFound() {
		return false
	}
	return true
}

func ChangeLevel(id int) error {
	if err := Db.Self.Model(&User{}).Where(&User{Id: id}).Update("level", 1).Error; err != nil {
		return err
	}
	return nil
}

func ChangeStatus(id int) error {
	var ischeckTmp Volunteer
	if err := Db.Self.Model(&Volunteer{}).Where(&Volunteer{Id: id}).Select("ischeck").Find(&ischeckTmp).Error; err != nil {
		return err
	}
	fmt.Println(ischeckTmp.IsCheck)
	if err := Db.Self.Model(&Volunteer{}).Where(&Volunteer{Id: id}).Update("ischeck", ischeckTmp.IsCheck+1).Error; err != nil {
		return err
	}
	return nil
}

func VerificationInfo(page, limit int) (l []CheckingInfo, err error) {

	if err := Db.Self.Table("volunteer").Model(&User{}).Not("ischeck", 5).Limit(limit).Offset((page - 1) * limit).Select("id,ischeck,name,phone").Scan(&l).Error; err != nil {
		return l, err
	}

	return l, nil
}
