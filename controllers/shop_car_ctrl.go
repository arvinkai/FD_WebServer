package controllers

import (
	"FD_WebServer/models"
	"strconv"

	"github.com/astaxie/beego"
)

type ShopCarController struct {
	beego.Controller
}

func (this *ShopCarController) Get() {
	strUid := this.Ctx.Request.Cookie("uid")
	Uid, _ := strconv.ParseInt(strUid, 10, 64)
	ShopcarsData := models.GetShopCarsData(Uid)
	this.TplName = "tab-subpage-shop-car.html"
}
