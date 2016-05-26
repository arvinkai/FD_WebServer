package controllers

import (
	"FD_WebServer/models"
	"fmt"

	"github.com/astaxie/beego"
)

type BookController struct {
	beego.Controller
}

func (this *BookController) Get() {
	//	str_uid, err := this.Ctx.Request.Cookie("uid")
	//	if err != nil {
	//		fmt.Println("BookController Get:", err)
	//	}
	//	uid, _ := strconv.ParseInt(str_uid.Value, 10, 64)
	//	booklist, count, err1 := models.GetBooklist(uid)
	booklist, count, err1 := models.GetBooklist(1)
	if count == 0 || err1 != nil {
		fmt.Println("GetBooklist count: ", count, " Error: ", err1)
	}

	this.Data["Booklist"] = booklist
	this.TplName = "tab-subpage-book.html"
}
