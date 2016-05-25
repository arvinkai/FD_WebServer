package controllers

import (
	"FD_WebServer/models"
	"fmt"

	"github.com/astaxie/beego"
)

type ShopController struct {
	beego.Controller
}

func (this *ShopController) Get() {
	Categorys, count, err := models.GetCategorys()
	if err != nil || count == 0 {
		this.Redirect("/", 301)
	}
	Goodsinfos, _ := models.GetGoodsinfoByCategory()

	fmt.Println(Goodsinfos)
	this.Data["Categorys"] = Categorys
	this.Data["Goodsinfos"] = Goodsinfos
	this.TplName = "tab-subpage-shop.html"
}
