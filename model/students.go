package model

// Struct holding the Student information
type Student struct {
	Binusian_ID  string `json:"binusian_id" xml:"binusian_id" form:"binusian_id" gorm:"primaryKey"`
	Fname        string `json:"fname" xml:"fname" form:"fname"`
	Lname        string `json:"lname" xml:"lname" form:"lname"`
	Email        string `json:"email" xml:"email" form:"email" gorm:"unique"`
	PasswordHash string `json:"password_hash" xml:"password_hash"`
	Phone_No     string `json:"phone_no" xml:"phone_no" form:"phone_no" gorm:"unique"`
	Description  string `json:"description" xml:"description" form:"description"`
	Profile_Img  string `json:"profile_img" xml:"profile_img" form:"profile_img"`
	Class_ID     string `json:"class_id" xml:"class_id" form:"class_id" gorm:"size:191"`
	Class        Class  `gorm:"foreignKey:Class_ID"`
}
