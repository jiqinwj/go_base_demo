package servicesb

import "go_base_demo/com.kris/core"

type NewsService struct  {

}

func (ns *NewsService) Get(id int) string {
	return "领导B-单新闻内容"
}


func  init()  {
	core.SetService(NewServiceFactory().Create("news"))
}