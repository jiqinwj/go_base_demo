package main

import (
	"net"
	"fmt"
)
func main()  {
	conn,err:=net.Dial("tcp","127.0.0.1:8099")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("i am shenyi")) //第一步

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(string(buf[:n]))






}