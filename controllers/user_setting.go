package controllers

import (
"github.com/astaxie/beego"
"FD_WebServer/models")

type UserSettingController struct {
	beego.Controller
}

func (this *UserSettingController) Get() {
	this.Data["ServerHost"] = beego.AppConfig.String("ServerHost")
	this.Data["ServerPort"] = beego.AppConfig.String("ServerPort")
    this.Data["Character"] =
	this.TplName = "tab-subpage-user.html"
}
