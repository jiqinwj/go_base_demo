package core

type NewsController struct {
	MyController
}

func(this *NewsController) GET()  {
   this.Ctx.WriteString("this is newscontroller")
}

func(this *NewsController) POST()  {
	this.Ctx.WriteString("this is newscontroller for POST")
}
