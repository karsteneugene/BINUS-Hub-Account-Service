package model

// Struct holding the class information
type Class struct {
	ID         string `json:"id" gorm:"primaryKey"`
	Class_Name string `json:"class_name" xml:"class_name" form:"class_name"`
}
