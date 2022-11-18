package model

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestStudentModel(t *testing.T) {
	validateStudent := validator.New()

	student := Student{
		Binusian_ID: "2440035596", Fname: "Darren", Lname: "Pangesa", Email: "darren.pangesa@binus.ac.id", PasswordHash: "", Phone_No: "081219201007", Description: "Lorem ipsum", Profile_Img: "<path of image>",
	}

	errStudent := validateStudent.Struct(student)

	if errStudent != nil {
		t.Error(errStudent.Error())
	}

}
