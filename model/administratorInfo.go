package model

type CheckingInfo struct {
	Id      int    `gorm:"column:id" ,json:"id"`
	IsCheck int    `gorm:"column:ischeck" ,json:"status"`
	Name    string `gorm:"column:name" ,json:"name"`
	Phone   string `gorm:"column:phone" ,json:"phone"`
}

type PhoneInfo struct {
	Phone string `json:"phone"`
}
