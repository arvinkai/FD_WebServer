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
	sliders, slidersCount, err := models.GetPicturesByUse("homehead", "slider", 1)
	if err != nil {
		fmt.Println(slidersCount, err)
	}
	this.Data["Pictures"] = sliders
	length := len(sliders)
	this.Data["PicturesLen"] = length - 1
	this.Data["StartPic"] = sliders[0]
	this.Data["EndPic"] = sliders[length-1]

	carousels, carouselCount, err1 := models.GetPicturesByUse("homecarousel", "carousel", 1)
	if err != nil {
		fmt.Println(carouselCount, err1)
	}
	this.Data["Carousels"] = carousels
	fmt.Println(carousels[0])
	showlists, showlistCount, err2 := models.GetPicturesByUse("homeshowlist", "showlist", 1)
	if err != nil {
		fmt.Println(showlistCount, err2)
	}
	this.Data["Showlists"] = showlists

	this.TplName = "tab-subpage-home.html"
}
