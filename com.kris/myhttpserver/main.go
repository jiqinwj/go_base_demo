package main

import (
	"net/http"

	"go_base_demo/com.kris/myhttpserver/core"
)

//type MyHandler struct {
//
//}
//func(*MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request){
//
//	 writer.Write([]byte("hello,myhandler"))
//}

func main()  {

	 router:=core.DefaultRouter()

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("get abc"))
	})
	router.Post("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("post abc"))
	})

	http.ListenAndServe(":8099",router)





}
