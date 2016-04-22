package controllers

import (
	"github.com/astaxie/beego"
)

type HomePageController struct {
	beego.Controller
}

func (this *HomePageController) Get() {
	this.TplName = "tab-subpage-home.html"
}
