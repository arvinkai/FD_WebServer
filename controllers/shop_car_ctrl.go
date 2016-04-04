package controllers

import (
	"github.com/astaxie/beego"
)

type ShopCarController struct {
	beego.Controller
}

func (this *ShopCarController) Get() {
	//beego.SetViewsPath("views/HomePage")
	this.TplName = "tab-subpage-shop-car.html"
}
