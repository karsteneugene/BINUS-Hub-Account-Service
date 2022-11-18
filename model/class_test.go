package model

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestClassModel(t *testing.T) {
	validateClass := validator.New()

	class := Class{
		ID: "L5AC", Class_Desc: "L4AC",
	}

	errClass := validateClass.Struct(class)

	if errClass != nil {
		t.Error(errClass.Error())
	}

}
