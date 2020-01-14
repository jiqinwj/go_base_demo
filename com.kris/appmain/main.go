package main

import (
	"fmt"
	"go_base_demo/com.kris/services"
)

func main()  {

   var service services.IService=services.NewServiceFactory().Create("user")

	fmt.Print(service.Get(123))

 


}


