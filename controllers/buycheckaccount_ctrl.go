// buycheckaccount_ctrl
package controllers

import (
	"github.com/astaxie/beego"
)

type BuyCheckAccountController struct {
	beego.Controller
}

func (this *BuyCheckAccountController) Get() {
	this.TplName = "buypage/buy_check_account.html"
}
