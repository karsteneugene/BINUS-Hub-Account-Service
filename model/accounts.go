package model

// Struct holding the account information
type Account struct {
	ID           uint   `json:"id" gorm:"column:id;PRIMARY_KEY"`
	Binusian_ID  string `json:"binusian_id"`
	Fname        string `json:"first_name"`
	Lname        string `json:"last_name"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash string `json:"password_hash"`
	Phone_No     string `json:"phone_no" gorm:"unique"`
	Role         string `json:"role"`
	Description  string `json:"description"`
	Profile_Img  string `json:"profile_img"`
}

// Struct holding the course information
type Course struct {
	ID          uint   `json:"id" gorm:"column:id;PRIMARY_KEY"`
	Course_Name string `json:"course_name"`
}

// Struct holding the class information
type Class struct {
	ID         uint   `json:"id" gorm:"column:id;PRIMARY_KEY"`
	Class_Name string `json:"class_name"`
}
