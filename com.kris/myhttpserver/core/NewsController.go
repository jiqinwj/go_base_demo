package core

import (

)

type NewsController struct {
	MyController
}

func(this *NewsController) GET()  {
   // this.Ctx.WriteString("this is newscontroller")
   p:=this.GetParam("username","no param","abc")
   this.Ctx.WriteString(p)

}

func(this *NewsController) POST()  {
	//this.Ctx.WriteString("this is newscontroller for POST")
	user:=UserMode{}
	this.JSONParam(&user)
	this.Ctx.WriteJSON(user)

}
