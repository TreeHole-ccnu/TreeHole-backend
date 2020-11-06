package model

import (
	//"github.com/TreeHole-ccnu/TreeHole-backend/util"

)

func ComfirmUserPhone(phone string, password string) (int, error) {
	var realpassword string

	if Db.Self.Model(&User{}).Where(&User{Phone:phone}).RecordNotFound() {
		return 1, nil
	}

	if err := Db.Self.Model(&User{}).Where(&User{Phone:phone}).Pluck("password", &realpassword).Error; err != nil {
		return 0, err
	}
	if password != realpassword {
		return 2, nil
	} else {
		return 3, nil
	}
}