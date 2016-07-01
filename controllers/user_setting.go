package controllers

import (
	"FD_WebServer/models"
	"fmt"

	"github.com/astaxie/beego"
)

type UserSettingController struct {
	beego.Controller
}

func (this *UserSettingController) Get() {
	this.TplName = "tab-subpage-user.html"
}

func (this *UserSettingController) Addonline() {
	uname := this.Input().Get("uname")
	token := this.Input().Get("token")

	userInfo, rel := models.GetUserInfo(uname, token)

	if rel != 0 {
		fmt.Println("UserOnline result:", rel)
	}
	onlineuert := &models.Online{Uid: userInfo.Uid, Uname: userInfo.Uname, Type: userInfo.Type, Statues: 1, Token: userInfo.Token}
	models.OperatOnline(onlineuert, "add")
}

func (this *UserSettingController) Delonline() {
	uname := this.Input().Get("uname")
	token := this.Input().Get("token")

	onlineUser, rel := models.GetOnlineUser(uname, token)

	if rel != 0 || onlineUser == nil {
		fmt.Println("UserOnline result:", rel)
	}

	models.OperatOnline(onlineUser, "del")
}
