package model

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestAccountsModel(t *testing.T) {
	validateAccount := validator.New()
	validateCourse := validator.New()
	validateClass := validator.New()

	account := Student{
		Binusian_ID: "2440035596", Fname: "Darren", Lname: "Pangesa", Email: "darren.pangesa@binus.ac.id", PasswordHash: "", Phone_No: "081219201007", Role: "student", Description: "Lorem ipsum", Profile_Img: "<path of image>",
	}

	course := Course{
		Course_ID: "0", Course_Name: "Distributed System", Class_ID: "L1AC", Lecturer_ID: "D6059",
	}

	class := Class{
		ID: 0, Class_Name: "L4AC",
	}

	errAccount := validateAccount.Struct(account)
	errCourse := validateCourse.Struct(course)
	errClass := validateClass.Struct(class)

	if errAccount != nil {
		t.Error(errAccount.Error())
	}

	if errCourse != nil {
		t.Error(errCourse.Error())
	}

	if errClass != nil {
		t.Error(errClass.Error())
	}

}
