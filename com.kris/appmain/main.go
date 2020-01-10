package main

import (
	"fmt"
	"go_base_demo/com.kris/services"
)

func main()  {



	result:=services.GetUser()

	fmt.Print(result)
}