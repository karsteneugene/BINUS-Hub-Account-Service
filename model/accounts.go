package model

// Struct holding the account information
type Account struct {
	ID           uint   `json:"id" xml:"id" form:"id" gorm:"primaryKey"`
	Binusian_ID  string `json:"binusian_id" xml:"binusian_id" form:"name"`
	Fname        string `json:"first_name" xml:"fname" form:"fname"`
	Lname        string `json:"last_name" xml:"lname" form:"lname"`
	Email        string `json:"email" xml:"email" form:"email" gorm:"unique"`
	PasswordHash string `json:"password_hash" xml:"password_hash"`
	Phone_No     string `json:"phone_no" xml:"phone_no" form:"phone_no" gorm:"unique"`
	Role         string `json:"role" xml:"role" form:"role"`
	Description  string `json:"description" xml:"description" form:"description"`
	Profile_Img  string `json:"profile_img" xml:"profile_img" form:"profile_img"`
	Course_Id    uint   `json:"course_id" xml:"course_id" form:"course_id"`
	Course       Course `gorm:"foreignKey:Course_Id"`
}

// Struct holding the course information
type Course struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Course_Name string `json:"course_name" xml:"course_name" form:"course_name"`
	Class_ID    uint   `json:"class_id" xml:"class_id" form:"class_id"`
	Class       Class  `gorm:"foreignKey:Class_Id"`
}

// Struct holding the class information
type Class struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Class_Name string `json:"class_name" xml:"class_name" form:"class_name"`
}
