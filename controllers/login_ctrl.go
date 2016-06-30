// Login_ctrl.go
package controllers

import (
	"FD_WebServer/models"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.Data["ServerHost"] = beego.AppConfig.String("ServerHost")
	this.Data["ServerPort"] = beego.AppConfig.String("ServerPort")
	this.TplName = "login/login.html"
}

func (this *LoginController) Userlogin() {
	uname := this.Input().Get("uname")
	pw := this.Input().Get("password")
	userInfo := models.UserLogin(uname, pw)
	result := 0
	if userInfo == nil {
		result = 1
	}

	if pw != userInfo.Pw {
		result = 2
	}

	this.Data["json"] = map[string]interface{}{"rel": result, "uname": userInfo.Uname, "nick": userInfo.Nickname,
		"qq": userInfo.Qqnum, "phone": userInfo.Phone, "email": userInfo.Email, "token": userInfo.Token}
	this.ServeJSON()
}
