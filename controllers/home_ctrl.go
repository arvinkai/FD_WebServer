package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	//	beego.SetViewsPath("views/OrderPage")
	this.TplName = "home.html"
}
