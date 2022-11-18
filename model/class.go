package model

// Struct holding the class information
type Class struct {
	ID         string `json:"id" xml:"id" form:"id" gorm:"primaryKey"`
	Class_Desc string `json:"class_desc" xml:"class_desc" form:"class_desc"`
}
