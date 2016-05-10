package controllers

import (
	"FD_WebServer/models"
	"fmt"

	"github.com/astaxie/beego"
)

type HomePageController struct {
	beego.Controller
}

func (this *HomePageController) Get() {
	pictures := models.GetPicture("Home")
	this.Data["Pictures"] = pictures
	this.Data["Picture0"] = pictures[0]
	this.Data["Picture2"] = pictures[1]
	fmt.Println(pictures[0])
	this.TplName = "tab-subpage-home.html"
}
