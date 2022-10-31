package model

// Struct holding the course information
type Course struct {
	Course_ID   string   `json:"id" gorm:"primaryKey"`
	Course_Name string   `json:"course_name" xml:"course_name" form:"course_name"`
	Class_ID    string   `json:"class_id" xml:"class_id" form:"class_id"`
	Lecturer_ID string   `json:"lecturer_id" xml:"lecturer_id" form:"lecturer_id"`
	Class       Class    `gorm:"foreignKey:Class_ID"`
	Lecturer    Lecturer `gorm:"foreignKey:Lecturer_ID"`
}
