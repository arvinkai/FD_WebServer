package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//	beego.SetViewsPath("views/HomePage")
	c.TplName = "tab-subpage-home.html"
}
