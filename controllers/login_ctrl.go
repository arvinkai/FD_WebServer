// Login_ctrl.go
package controllers

import (
	"FD_WebServer/models"
	"fmt"

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

func (this *LoginController) Addonline() {
	uname := this.Input().Get("uname")
	token := this.Input().Get("token")

	userInfo, rel := models.GetUserInfo(uname, token)

	if rel != 0 || userInfo == nil {
		fmt.Println("UserOnline result:", rel)

	} else {
		onlineuert := &models.Online{Uid: userInfo.Uid, Uname: userInfo.Uname, Type: userInfo.Type, Statues: 1, Token: userInfo.Token}
		models.OperatOnline(onlineuert, "add")
	}

	this.Data["json"] = map[string]interface{}{"rel": rel}
	this.ServeJSON()
}

func (this *LoginController) Delonline() {
	uname := this.Input().Get("uname")
	token := this.Input().Get("token")

	onlineUser, rel := models.GetOnlineUser(uname, token)

	if rel != 0 || onlineUser == nil {
		fmt.Println("UserOnline result:", rel)
	} else {
		models.OperatOnline(onlineUser, "del")
	}

	this.Data["json"] = map[string]interface{}{"rel": rel}
	this.ServeJSON()
}
