package controllers

import (
	"github.com/astaxie/beego"
)

type ShopController struct {
	beego.Controller
}

func (this *ShopController) Get() {
	this.TplName = "tab-subpage-shop.html"
}
