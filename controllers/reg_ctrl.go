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
	fmt.Println("uname:", uname)
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
	this.Data["json"] = map[string]interface{}{"UserName": uname, "CanUse": canUse}
	this.ServeJSON()
}

func (this *RegistController) RegInfo() {

}
