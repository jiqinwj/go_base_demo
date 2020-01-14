package services

type UserService struct {

}

func (us *UserService) Get(id int) string  {
	return "单个用户信息"
}