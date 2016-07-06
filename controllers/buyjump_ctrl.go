// buyjump_ctrl.go
package controllers

import (
	"FD_WebServer/models"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type BuyJumpController struct {
	beego.Controller
}

func (this *BuyJumpController) Get() {
	this.Addtocar()
	this.Data["ServerHost"] = beego.AppConfig.String("ServerHost")
	this.Data["ServerPort"] = beego.AppConfig.String("ServerPort")
	this.TplName = "buypage/buy_jump_page.html"
}

func (this *BuyJumpController) Addtocar() {
	//	var uid int64 = 0
	//	ck_uid, err := this.Ctx.Request.Cookie("uid")
	//	if err != nil {
	//		fmt.Println(err)
	//	} else {
	//		uid, _ = strconv.ParseInt(ck_uid, 10, 64)
	//	}
	str_goodsid := this.Input().Get("goodsid")
	goodsid, _ := strconv.ParseInt(str_goodsid, 10, 64)
	count, _ := strconv.ParseInt(this.Input().Get("count"), 10, 32)
	if goodsid == 0 || count == 0 {
		fmt.Println(goodsid, count)
		return
	}
	v := time.Now().Unix()
	nowDate := strconv.FormatInt(v, 10)
	//	spcar := &models.Shopcar{Uid: uid, Goodsid: goodsid, Count: count, CreateDate: nowDate}
	spcar := &models.Shopcar{Uid: 1, Goodsid: goodsid, Count: int32(count), CreateDate: nowDate}
	count, err := models.OpShopcar(spcar, "add")
	if err != nil {
		fmt.Println(err)
	}
}
