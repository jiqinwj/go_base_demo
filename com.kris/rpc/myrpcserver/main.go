package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type UserService struct {

}

func(this *UserService) GetUserName(userid int,username *string) error  {
   if userid==101{
   	*username="shenyi"
   }else{
	   *username="guest"
   }
	return nil
}
func main(){

	lis,_:=net.Listen("tcp",":8082")
	rpc.Register(new(UserService))//注册对外暴露的方法

	for{
		client,_:=lis.Accept() //很常规的socket 编程哦~~~
		jsonrpc.ServeConn(client) //处理客户端连接
	}






}
