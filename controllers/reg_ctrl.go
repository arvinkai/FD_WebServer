// reg_ctrl.go
package controllers

import (
	"FD_WebServer/models"

	"github.com/astaxie/beego"
)

type RegistController struct {
	beego.Controller
}

func (this *RegistController) Get() {
	this.Data["ServerHost"] = beego.AppConfig.String("ServerHost")
	this.Data["ServerPort"] = beego.AppConfig.String("ServerPort")
	this.Data["CanUse"] = false

	this.TplName = "login/reg.html"
}

func (this *RegistController) Checkuname() {
	uname := this.Input().Get("uName")
	rel, _ := models.CheckUserName(uname)

	this.Data["json"] = map[string]interface{}{"UserName": uname, "CanUse": rel}
	this.ServeJSON()
}

func (this *RegistController) RegInfo() {
	uname := this.Input().Get("uname")
	pw := this.Input().Get("pw")
	nickname := this.Input().Get("nickname")
	email := this.Input().Get("email")
	telNum := this.Input().Get("telNum")

	uid, rel, err := models.CreateUser(uname, pw, nickname, email, telNum)
	success := true
	if err != nil || uid == -1 || !rel {
		success = false
	}
	this.Data["json"] = map[string]interface{}{"result": success}
	this.ServeJSON()
}
