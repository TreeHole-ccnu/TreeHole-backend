package model

type User struct {
	Phone          string `gorm:"column:phone", json:"phone"`
	Password       string `gorm:"colmun:password", json:"password"`
	Id             int    `gorm:"column:id", gorm:"AUTO_INCREMENT", json:"-"'`
	Level          int    `gorm:"column:level"`
	Name           string `gorm:"column:name", json:"name"`
	Sex            string `gorm:"column:sex", json:"sex"`
	Birth          string `gorm:"column:birth", json:"birth"`
	Nation         string `gorm:"column:nation", json:"nation"`
	NativePlace    string `gorm:"column:native_place", json:"native_place"`
	Email          string `gorm:"column:email", json:"email"`
	IdentityNumber string `gorm:"column:identity_number",json:"identity_number"`
	ImageUrl       string `gorm:"column:image_url", json:"image_url"`
}

type LoginInfo struct {
	Phone    string `gorm:"column:phone", json:"phone"`
	Password string `gorm:"colmun:password", json:"password"`
}

type RegisterInfo struct {
	Phone    string `gorm:"column:phone", json:"phone"`
	Vcd      string `gorm:"column:vcd", json:"vcd"`
	Password string `gorm:"column:password", json:"password"`
}
