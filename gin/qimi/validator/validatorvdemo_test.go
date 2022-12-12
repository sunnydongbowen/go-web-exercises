package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

type DATA struct {
	Email string `validate:"required,email"`
	Phone string `validate:"required"`
}

func f1() {
	// 声明一个校验器对象
	v := validator.New()
	su := DATA{
		Email: "xxx",
	}
	err := v.Struct(&su)
	fmt.Println(err)
}

func TestValidatordemo(t *testing.T) {
	f1()
}
