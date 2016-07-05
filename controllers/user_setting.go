package controllers

import "github.com/astaxie/beego"

type UserSettingController struct {
	beego.Controller
}

func (this *UserSettingController) Get() {
	this.Data["ServerHost"] = beego.AppConfig.String("ServerHost")
	this.Data["ServerPort"] = beego.AppConfig.String("ServerPort")
	this.TplName = "tab-subpage-user.html"
}
