package models

type UserModel struct {
	Uid   int
	Uname string
}

func (u *UserModel) SetValue(id int, name string) {
	u.Uname = name
	u.Uid = id
}

func (u *UserModel) ToString() string {
	return "用户名是" + u.Uname
}

/*
class UserModel
{
  public String ToString()
{
   //////
}

}
*/
