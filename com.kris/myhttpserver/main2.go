package main

//type MyHandler struct {
//
//}
//func(*MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request){
//
//	 writer.Write([]byte("hello,myhandler"))
//}
//
//func main()  {
//
//	 mymux:=http.NewServeMux()
//	 mymux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
//		 writer.Write([]byte("index"))
//	 })
//	 mymux.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
//		getUserName:= request.URL.Query().Get("username")
//		if getUserName!=""{
//			c:=&http.Cookie{Name:"username",Value:getUserName,Path:"/"}
//			http.SetCookie(writer,c)
//		}
//		writer.Write([]byte("这里是登录页面"));
//	 })
//	 mymux.HandleFunc("/unlogin", func(writer http.ResponseWriter, request *http.Request) {
//		 c:=&http.Cookie{Name:"username",Value:"",Path:"/",Expires:time.Now().AddDate(-1,0,0)}
//		 http.SetCookie(writer,c)
//		 writer.Header().Set("Content-type","text/html")
//		 writer.Write([]byte("unlogin...."));
//		 script:=`<script>
//			setTimeout(()=>{self.location='/'},2000)
//
//</script>`
//		 writer.Write([]byte(script))
//	 })
//
//	http.ListenAndServe(":8099",mymux)
//
//
//
//
//
//}
