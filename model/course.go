package model

// Struct holding the course information
type Course struct {
	ID          uint     `json:"id" xml:"id" form:"id" gorm:"primaryKey"`
	Course_ID   string   `json:"course_id" xml:"course_id" form:"course_id"`
	Course_Name string   `json:"course_name" xml:"course_name" form:"course_name"`
	Class_ID    string   `json:"class_id" xml:"class_id" form:"class_id" gorm:"size:191"`
	Lecturer_ID string   `json:"lecturer_id" xml:"lecturer_id" form:"lecturer_id" gorm:"size:191"`
	Class       Class    `gorm:"foreignKey:Class_ID"`
	Lecturer    Lecturer `gorm:"foreignKey:Lecturer_ID"`
}
