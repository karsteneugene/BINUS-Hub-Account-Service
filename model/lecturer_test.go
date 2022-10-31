package model

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestLecturerModel(t *testing.T) {
	validateLecturer := validator.New()

	lecturer := Lecturer{
		Lecturer_ID: "2440035596", Fname: "Darren", Lname: "Pangesa", Email: "darren.pangesa@binus.ac.id", PasswordHash: "", Phone_No: "081219201007", Role: "student", Description: "Lorem ipsum", Profile_Img: "<path of image>",
	}

	errLecturer := validateLecturer.Struct(lecturer)

	if errLecturer != nil {
		t.Error(errLecturer.Error())
	}

}
