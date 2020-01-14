package main

import (
	"fmt"
	"go_base_demo/com.kris/models"
)

func main()  {



	//result:=services.GetUser()

	news:=models.UserModel{123,"title"}
	fmt.Println(news.ToJSON())

	//
	//fmt.Print(result)
}