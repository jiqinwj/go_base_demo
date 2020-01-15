package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

func main()  {
  client,eror:= jsonrpc.Dial("tcp","127.0.0.1:8082")
  if eror!=nil{
  	fmt.Println(eror)
  	return
  }
  username:=""
  err:=client.Call("UserService.GetUserName", 101,&username)
  if err!=nil{
  	fmt.Println(err)
  	return
  }
  fmt.Println(username)

}