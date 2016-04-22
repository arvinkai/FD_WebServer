package controllers

import (
	"github.com/astaxie/beego"
)

type BookController struct {
	beego.Controller
}

func (this *BookController) Get() {
	this.TplName = "tab-subpage-book.html"
}
