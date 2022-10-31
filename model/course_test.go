package model

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestCourseModel(t *testing.T) {
	validateCourse := validator.New()

	course := Course{
		Course_ID: "0", Course_Name: "Distributed System", Class_ID: "L1AC", Lecturer_ID: "D6059",
	}

	errCourse := validateCourse.Struct(course)

	if errCourse != nil {
		t.Error(errCourse.Error())
	}

}
