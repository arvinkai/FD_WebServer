package controllers

import (
	"github.com/astaxie/beego"
)

type ShopCarController struct {
	beego.Controller
}

func (this *ShopCarController) Get() {
	this.TplName = "tab-subpage-shop-car.html"
}
