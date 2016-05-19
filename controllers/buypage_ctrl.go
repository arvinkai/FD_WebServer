// buypage_ctrl.go
package controllers

import (
	"FD_WebServer/models"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type BuyPageController struct {
	beego.Controller
}

func (this *BuyPageController) Get() {
	this.TplName = "buypage/buy_item_page.html"
}

func (this *BuyPageController) Addtocar() {

	v := time.Now().Unix()
	nowDate := strconv.FormatInt(v, 10)
	spcar := &models.Shopcar{Uid: 1, Goodsid: 3, Count: 2, CreateDate: nowDate}
	err := models.OpShopcar(spcar, "add")
	if err != nil {

	}
	this.TplName = "buypage/buy_item_page.html"
}
