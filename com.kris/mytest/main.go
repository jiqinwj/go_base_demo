package main

import (
	"regexp"
	"fmt"
)

func GetRqPath(rq string) string  {
	r:=regexp.MustCompile(`^GET\s(.*?)\sHTTP`)
	if r.MatchString(rq) {
		return r.FindStringSubmatch(rq)[1]
	} else{
		return "/"
	}
}
func main()  {

	str:=`GET /abc.php HTTP/1.1
Host: localhost:8099
Connection: keep-alive
Upgrade-Insecure-Requests: 1`
	fmt.Println(GetRqPath(str))
}
