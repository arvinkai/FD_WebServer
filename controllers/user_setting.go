package controllers

import (
	"github.com/astaxie/beego"
)

type UserSettingController struct {
	beego.Controller
}

func (this *UserSettingController) Get() {
	this.TplName = "tab-subpage-shop.html"
}
