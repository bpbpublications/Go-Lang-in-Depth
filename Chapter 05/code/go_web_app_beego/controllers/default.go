package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (main *MainController) Get() {
	main.Data["Website"] = "beego.goapp"
	main.Data["Email"] = "gouser@gmail.com"
	main.TplName = "index.tpl"
}

func (main *MainController) BeegoWebApp() {
	main.Data["Website"] = "beego webapp"
	main.Data["Email"] = "gouser@example.com"
	main.Data["EmailName"] = "Go User"
	main.TplName = "default/hello-beego.tpl"
}
