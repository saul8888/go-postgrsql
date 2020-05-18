package validate

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	ID        string `validate:"omitempty,uuid"`
	Firstname string `validate:"required"`
	Lastname  string `validate:"required"`
	//Age       int    `validate:"gte=0,lte=80"`
	//Field    string `validate:"excludesall=,"` // BAD! Do not include a comma.
	Username string `validate:"required,email"`
	Password string `validate:"required,gte=6,lte=15"`
	Type     string `validate:"isdefault"`
}

func main() {
	user := User{
		Firstname: "Saul",
		Lastname:  "Quispe",
		Username:  "saul@example.com",
		Password:  "1234567890",
		//Type:      "something",
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}
