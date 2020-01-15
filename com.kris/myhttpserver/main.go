package main

import (
	"go_base_demo/com.kris/myhttpserver/core"
	"net/http"
)

//type MyHandler struct {
//
//}
//func(*MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request){
//
//	 writer.Write([]byte("hello,myhandler"))
//}

func main()  {



	//router:=core.DefaultRouter()
	//
	//router.Get("/", func(ctx *core.MyContext) {
	//	ctx.WriteString("my string GET")
	//})
	//router.Post("/", func(ctx *core.MyContext) {
	//	ctx.WriteString("my string POST")
	//})
	//
	//http.ListenAndServe(":8099",router)


	//TODO 上下文版本
	//
	//
	//router:=core.DefaultRouter()
	//
	//router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
	//		writer.Write([]byte("get abc"))
	//})
	//router.Post("/", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("post abc"))
	//})
	//
	//http.ListenAndServe(":8099",router)
	//

   //TODO 控制器版本
	router:=core.DefaultRouter()

	router.Add("/",&core.NewsController{})


	http.ListenAndServe(":8099",router)




}
