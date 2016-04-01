package controllers

import (
	"github.com/astaxie/beego"
)

type OrderController struct {
	beego.Controller
}

func (this *OrderController) Get() {
	//	beego.SetViewsPath("views/OrderPage")
	this.TplName = "order.html"
}
