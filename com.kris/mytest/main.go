package main

import (
	"context"
	"net/http"
	"time"
)

//TODO 正则乱入

//func GetRqPath(rq string) string  {
//	r:=regexp.MustCompile(`^GET\s(.*?)\sHTTP`)
//	if r.MatchString(rq) {
//		return r.FindStringSubmatch(rq)[1]
//	} else{
//		return "/"
//	}
//}
//func main()  {
//
//	str:=`GET /abc.php HTTP/1.1
//Host: localhost:8099
//Connection: keep-alive
//Upgrade-Insecure-Requests: 1`
//	fmt.Println(GetRqPath(str))
//}


////TODO context 协程之间信号 传递
//func main()  {
//	rand.Seed(time.Now().Unix())
//
//	ctx,_:=context.WithTimeout(context.Background(),time.Second*3)
//
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go GenUsers(ctx,&wg)
//	wg.Wait()
//
//	fmt.Println("生成幸运用户成功")
//}
//func GenUsers(ctx context.Context,wg *sync.WaitGroup)  { //生成用户ID
//	fmt.Println("开始生成幸运用户")
//	users:=make([]int,0)
//guser:for{
//	select{
//	case <- ctx.Done(): //代表父context发起 取消操作
//
//		fmt.Println(users)
//		wg.Done()
//		break guser
//		return
//	default:
//		users=append(users,getUserID(1000,100000))
//	}
//}
//
//}
//func getUserID(min int ,max int) int  {
//	return rand.Intn(max-min)+min
//}


//TODO httpserver 超时操作
func CountData(c chan string) chan string {
	time.Sleep(time.Second*8)
	c<- "统计结果"
	return c
}

type IndexHandler struct {}
func(this *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Query().Get("count")==""{
		w.Write([]byte("这是首页"))
	}else {
		ctx,cancel:=context.WithTimeout(r.Context(),time.Second*3)
		defer cancel()
		c:=make(chan string)
		go CountData(c)
		select {
		case <-ctx.Done():
			w.Write([]byte("超时"))
		case ret:=<-c:
			w.Write([]byte(ret))
		}


	}


}

func main()  {
	mux:=http.NewServeMux()
	mux.Handle("/",new(IndexHandler))

	http.ListenAndServe(":8082",mux)
}




