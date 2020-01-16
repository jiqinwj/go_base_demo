package core

type IService interface {
	Get(id int) string
}

var myService IService

func SetService(service IService) {
	myService = service
}
func GetService() IService {
	return myService
}
