// reg_ctrl.go
package controllers

import (
	"FD_WebServer/models"
	"fmt"

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
	fmt.Println(uname)
	rel, err := models.CheckUserName(uname)
	var canUse bool = true
	if err != nil {
		canUse = false
	} else {

		if !rel {
			canUse = false
		} else {
			canUse = true
		}
	}
	this.Data["json"] = map[string]interface{}{"UserName": uname, "CanUse": canUse, "ServerHost": 123}
	this.Data["ServerHost"] = beego.AppConfig.String("ServerHost")
	this.Data["ServerPort"] = beego.AppConfig.String("ServerPort")
	this.ServeJSON()
	//	this.Data["UserName"] = uname
	//	this.Data["CanUse"] = canUse
	//	this.Data["ServerHost"] = beego.AppConfig.String("ServerHost")
	//	this.Data["ServerPort"] = beego.AppConfig.String("ServerPort")
	fmt.Println(canUse)
	//	this.Redirect("/home", 302)

	//	this.TplName = "login/reg.html"
}
