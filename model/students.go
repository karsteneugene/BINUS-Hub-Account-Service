package model

// Struct holding the Student information
type Student struct {
	Binusian_ID  string `json:"binusian_id" xml:"binusian_id" form:"binusian_id" gorm:"primaryKey"`
	Fname        string `json:"first_name" xml:"fname" form:"fname"`
	Lname        string `json:"last_name" xml:"lname" form:"lname"`
	Email        string `json:"email" xml:"email" form:"email" gorm:"unique"`
	PasswordHash string `json:"password_hash" xml:"password_hash"`
	Phone_No     string `json:"phone_no" xml:"phone_no" form:"phone_no" gorm:"unique"`
	Role         string `json:"role" xml:"role" form:"role"`
	Description  string `json:"description" xml:"description" form:"description"`
	Profile_Img  string `json:"profile_img" xml:"profile_img" form:"profile_img"`
	Course_Id    uint   `json:"course_id" xml:"course_id" form:"course_id"`
	Course       Course `gorm:"foreignKey:Course_ID"`
}
