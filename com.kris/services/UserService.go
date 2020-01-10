package services

import (
	"go_base_demo/com.kris/models"
)

func GetUser()  string {

    user:=new(models.UserModel)
    user.SetValue(123,"lisi")

	return user.ToString()
}

