package main

import (
	"net/rpc"
	"fmt"
)

func main()  {
  client,_:= rpc.Dial("tcp","127.0.0.1:8082")
  username:=""
  err:=client.Call("UserService.GetUserName",102,&username)
  if err!=nil{
  	fmt.Println(err)
  	return
  }
  fmt.Println(username)

}