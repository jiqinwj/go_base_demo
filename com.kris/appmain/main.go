package main

import (
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
	"sync"
	"time"
)

func main() {

	//TODO 数据库练习 切片 扩展运算 等
	//db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8mb4")
	//if err != nil {
	//	fmt.Println("连接错" + err.Error())
	//	return
	//}
	//rows, err := db.Query("select user_id,user_name from users1 limit 0,10")
	//if err != nil {
	//	fmt.Println("查询错误" + err.Error())
	//	return
	//}
	//
	//columns, _ := rows.Columns()
	//allRows := make([]interface{}, 0) //所有行  大切片
	//fieldMap := make(map[string]string)
	//oneRow := make([]interface{}, len(columns)) //定义一行切片
	//scanRow := make([]interface{}, len(columns))
	//for rows.Next() {
	//
	//	for i, _ := range oneRow {
	//
	//		scanRow[i] = &oneRow[i]
	//	}
	//	rows.Scan(scanRow...)
	//
	//	for i, val := range oneRow {
	//		v, ok := val.([]byte) //断言
	//		if (ok) {
	//			fieldMap[columns[i]] = string(v)
	//		}
	//	}
	//	allRows = append(allRows, fieldMap)
	//}
	//fmt.Println(allRows)

	//TODO 协程练习
	//num:=100000000
	//start:=time.Now()
	//c:=make(chan int,2)
	//go sum(1,num/2,c);
	//go sum(num/2+1,num,c)
	//
	//ret1,ret2:=<- c,<- c
	//
	//end:=time.Now()
	//fmt.Println(end.Sub(start),ret1+ret2)

	//TODO 协程练习2
	//users := strings.Split("zhangsan,ss,lisi,wangwu", ",")
	//ages := strings.Split("19,21,25,26", ",")
	//
	//fmt.Println(ages)
	//
	//c1, c2 := make(chan bool), make(chan bool)
	//ret := make([]string, 0)
	//go func() {
	//	for _, v := range users {
	//		<-c1
	//		ret = append(ret, v)
	//		//time.Sleep(time.Second)
	//		c2 <- true
	//	}
	//}()
	//go func() {
	//	for _, v := range ages {
	//		<-c2
	//		ret = append(ret, v)
	//		c1 <- true
	//	}
	//}()
	//c1 <- true
	//fmt.Println(ret)

	//context:="xx"
	//
	//ioutil.WriteFile("./files/1.txt",[]byte(context),777)

	//TODO 抓取数据
	//url:="https://news.cnblogs.com/n/page/%d/"
	//c:=make(chan map[int][]byte)
	//for i:=1;i<=3;i++{
	//	go func(index int) {
	//		defer func() {
	//			if err:=recover();err!=nil{
	//				fmt.Println(err)
	//			}
	//			if index==3{
	//				close(c)
	//			}
	//		}()
	//		url:=fmt.Sprintf(url,index)  //格式化字符串，生成每一页的页码url
	//		res,_:=http.Get(url)  //利用http库 读取网页内容，返回response对象
	//		cnt,_:= ioutil.ReadAll(res.Body) //读出response的body部分
	//		if index==3{
	//			test()
	//		}
	//		c<-map[int][]byte{index:cnt} //把结果构建成一个map 插入到channel 里面
	//
	//	}(i)
	//}
	//
	//
	//for getcnt:=range c{ //主线程干的事，不断的range channel,直至channel被关闭
	//	for k,v:=range getcnt{
	//		ioutil.WriteFile(fmt.Sprintf("E:/go_start/go_base_demo/com.kris/files/%d",k),v,666)  //把内容写入到文件中
	//
	//	}
	//
	//}

	//	//TODO 协程阻塞
	//	url := "https://news.cnblogs.com/n/page/%d/"
	//	c := make(chan map[int][]byte)
	//	for i := 1; i <= 10; i++ {
	//		go func(index int) {
	//			defer func() {
	//				if err := recover(); err != nil {
	//					fmt.Println(err)
	//				}
	//
	//			}()
	//			url := fmt.Sprintf(url, index)     //格式化字符串，生成每一页的页码url
	//			res, _ := http.Get(url)            //利用http库 读取网页内容，返回response对象
	//			cnt, _ := ioutil.ReadAll(res.Body) //读出response的body部分
	//
	//			c <- map[int][]byte{index: cnt} //把结果构建成一个map 插入到channel 里面
	//		}(i)
	//	}
	//
	//	//
	//	//for getcnt:=range c{ //主线程干的事，不断的range channel,直至channel被关闭
	//	//
	//	//	 for k,v:=range getcnt{
	//	//	  ioutil.WriteFile(fmt.Sprintf("./files/%d",k),v,666)  //把内容写入到文件中
	//	//}
	//	//
	//	//
	//	//}
	//	result := map[int][]byte{}
	//myloop:
	//	for {
	//		select {
	//		case result = <-c:
	//			for k, v := range result {
	//				ioutil.WriteFile(fmt.Sprintf("E:/go_start/go_base_demo/com.kris/files/%d",k),v,666)  //把内容写入到文件中
	//
	//			}
	//		case <-time.After(time.Second * 3):
	//			break myloop
	//
	//		}
	//
	//	}

	//TODO sync 阻塞
	//var wg sync.WaitGroup
	//
	//url := "https://news.cnblogs.com/n/page/%d/"
	//// c:=make(chan map[int][]byte)
	//for i := 1; i <= 10; i++ {
	//	go func(index int) {
	//		defer func() {
	//			if err := recover(); err != nil {
	//				fmt.Println(err)
	//			}
	//			wg.Done()
	//		}()
	//		url := fmt.Sprintf(url, index)     //格式化字符串，生成每一页的页码url
	//		res, _ := http.Get(url)            //利用http库 读取网页内容，返回response对象
	//		cnt, _ := ioutil.ReadAll(res.Body) //读出response的body部分
	//		//c<-map[int][]byte{index:cnt} //把结果构建成一个map 插入到channel 里面
	//		ioutil.WriteFile(fmt.Sprintf("E:/go_start/go_base_demo/com.kris/files/%d", index), cnt, 666)
	//	}(i)
	//	wg.Add(1)
	//}
	////result:=map[int][]byte{}
	////myloop:for{
	////select {
	////case result=<-c:
	////	 for k,v:=range result{
	////		 ioutil.WriteFile(fmt.Sprintf("./files/%d",k),v,666)  //把内容写入到文件中
	////	 }
	////case <-time.After(time.Second*3):
	////	 	break myloop
	////
	////}
	//wg.Wait()
	//fmt.Println("抓取完毕")



	//TODO 协程waitGroup + 互斥锁
	var wg sync.WaitGroup
	var locker sync.Mutex
	file,_:=os.OpenFile("E:/go_start/go_base_demo/com.kris/test",os.O_RDONLY,666)
	defer  file.Close()
	fw:=bufio.NewReader(file)

	for i:=1;i<=2;i++{
		go func(index int) {
			defer  wg.Done()
			for{
				locker.Lock()
				str,err:=fw.ReadString('\n')
				if err!=nil{
					if err==io.EOF{
						locker.Unlock()
						break
					}
					fmt.Println(err)
				}
				time.Sleep(time.Millisecond*200)
				fmt.Printf("【协程%d】:%s",index,str)
				locker.Unlock()
			}
		}(i)
	}
	wg.Add(2)
	wg.Wait()
	fmt.Println("读取完成")





}

func test() {
	panic("exp")
}

func sum(min, max int, c chan int) {
	result := 0
	for i := min; i <= max; i++ {
		result = result + i
	}
	c <- result
}
