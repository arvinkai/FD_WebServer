package controllers

import (
	"github.com/astaxie/beego"
)

type ShopController struct {
	beego.Controller
}

func (this *ShopController) Get() {
	//	beego.SetViewsPath("views/ShopPage")
	this.TplName = "shop.html"
}
