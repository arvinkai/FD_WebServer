// checkaccount.go
package controllers

import "github.com/astaxie/beego"

type CheckAccountController struct {
	beego.Controller
}

func (this *CheckAccountController) Get() {
	this.TplName = "buypage/buy_check_account.html"
}
