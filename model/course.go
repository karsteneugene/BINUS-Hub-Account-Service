package model

// Struct holding the course information
type Course struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	Course_ID   string   `json:"course_id"`
	Course_Name string   `json:"course_name" xml:"course_name" form:"course_name"`
	Class_ID    string   `json:"class_id" xml:"class_id" form:"class_id"`
	Lecturer_ID string   `json:"lecturer_id" xml:"lecturer_id" form:"lecturer_id"`
	Class       Class    `json:"class" gorm:"foreignKey:Class_ID; references:id"`
	Lecturer    Lecturer `json:"lecturer" gorm:"foreignKey:Lecturer_ID; references:lecturer_id"`
}
