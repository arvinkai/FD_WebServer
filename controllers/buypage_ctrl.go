// buypage_ctrl.go
package controllers

import (
	"github.com/astaxie/beego"
)

type BuyPageController struct {
	beego.Controller
}

func (this *BuyPageController) Get() {
	this.TplName = "buypage/buy_item_page.html"
}
