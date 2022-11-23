package model

// Struct holding the lecturer information
type Lecturer struct {
	Lecturer_ID  string `json:"binusian_id" xml:"binusian_id" form:"binusian_id" gorm:"primaryKey"`
	Fname        string `json:"fname" xml:"fname" form:"fname"`
	Lname        string `json:"lname" xml:"lname" form:"lname"`
	Email        string `json:"email" xml:"email" form:"email" gorm:"unique"`
	PasswordHash string `json:"-" xml:"-" form:"-"`
	Phone_No     string `json:"phone_no" xml:"phone_no" form:"phone_no" gorm:"unique"`
	Description  string `json:"description" xml:"description" form:"description"`
	Profile_Img  string `json:"profile_img" xml:"profile_img" form:"profile_img"`
}
