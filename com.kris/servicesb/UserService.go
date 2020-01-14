package servicesb

type UserService struct {

}

func (us *UserService) Get(id int) string  {
	return "领导B-单个用户信息"
}
