package model

type VolunteerInfo struct {
	Education        string       `gorm:"column:education" json:"education"`
	Reference        string       `gorm:"column:reference" json:"reference"`
	Political        string       `gorm:"column:political" json:"political"`
	Nationality      string       `gorm:"column:nationality" json:"nationality"`
	Workphone        string       `gorm:"column:workphone" json:"workphone"`
	Job              string       `gorm:"column:job" json:"job"`
	SocialJob        string       `gorm:"column:social_job" json:"social_job"`
	MedicalHistory   string       `gorm:"column:medical_history" json:"medical_history"`
	TreatmentHistory string       `gorm:"column:treatment_history" json:"treatment_history"`
	Medicine         string       `gorm:"column:medicine" json:"medicine"`
	Reason           string       `gorm:"column:reason" json:"reason"`
	FrontUrl         string       `gorm:"column:front_url" json:"front_url"`
	ContraryUrl      string       `gorm:"column:contrary_url" json:"contrary_url"`
	ResumeInfos      []ResumeInfo `json:"resume"`
}

type Volunteer struct {
	Id               int    `gorm:"column:id" json:"-"`
	UserId           int    `gorm:"column:user_id" json:"user_id"`
	IsCheck          int    `gorm:"column:ischeck" json:"-"`
	Reference        string `gorm:"column:reference" json:"reference"`
	Name             string `gorm:"column:name" json:"name"`
	Birth            string `gorm:"column:birth" json:"birth"`
	Political        string `gorm:"column:political" json:"political"`
	Sex              string `gorm:"column:sex" json:"sex"`
	Nation           string `gorm:"column:nation" json:"nation"`
	NativePlace      string `gorm:"column:native_place" json:"native_place"`
	Education        string `gorm:"column:education" json:"education"`
	Nationality      string `gorm:"column:nationality" json:"nationality"`
	IdentityNumber   string `gorm:"column:identity_number" json:"identity_number"`
	WorkPhone        string `gorm:"column:workphone" json:"work_phone"`
	Phone            string `gorm:"column:phone" json:"phone"`
	Email            string `gorm:"column:email" json:"email"`
	Job              string `gorm:"column:job" json:"job"`
	SocialJob        string `gorm:"column:social_job" json:"social_job"`
	MedicalHistory   string `gorm:"column:medical_history" json:"medical_history"`
	TreatmentHistory string `gorm:"column:treatment_history" json:"treatment_history"`
	Medicine         string `gorm:"column:medicine" json:"medicine"`
	Reason           string `gorm:"column:reason" json:"reason"`
	FrontUrl         string `gorm:"column:front_url" json:"front_url"`
	ContraryUrl      string `gorm:"column:contrary_url" json:"contrary_url"`
	Time             string `gorm:"column:time" json:"-"`
}

type ResumeInfo struct {
	Date      string `gorm:"column:date" json:"date"`
	WorkPlace string `gorm:"column:work_place" json:"work_place"`
	Job       string `gorm:"column:job" json:"job"`
}

type Resume struct {
	Id        int    `gorm:"column:id"`
	Date      string `gorm:"column:date"`
	WorkPlace string `gorm:"column:work_place"`
	Job       string `gorm:"column:job"`
}
