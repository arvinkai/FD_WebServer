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
	pictures, clos, err := models.GetPicture("homehead", "slider", 1)
	if err != nil {
		fmt.Println(err)
		this.Redirect("/", 301)
	}
	this.Data["Pictures"] = pictures
	length := len(pictures)
	this.Data["PicturesLen"] = length - 1
	this.Data["StartPic"] = pictures[0]
	this.Data["EndPic"] = pictures[length-1]
	fmt.Println(pictures[0])
	this.TplName = "tab-subpage-home.html"
}
