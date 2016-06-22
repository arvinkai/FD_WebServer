// buypage_ctrl.go
package controllers

import (
	"FD_WebServer/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type BuyPageController struct {
	beego.Controller
}

func (this *BuyPageController) Get() {
	cl_Goodsid, _ := strconv.ParseInt(this.Input().Get("goodsid"), 10, 64)
	this.GetPageData(cl_Goodsid)
	this.Data["ServerHost"] = beego.AppConfig.String("ServerHost")
	this.Data["ServerPort"] = beego.AppConfig.String("ServerPort")

	this.TplName = "buypage/buy_item_page.html"
}

func (this *BuyPageController) Test() {
	fmt.Println("test")
	this.TplName = "buypage/buy_item_page.html"
}

//func (this *BuyPageController) Addtocar() {
//	//	var uid int64 = 0
//	//	ck_uid, err := this.Ctx.Request.Cookie("uid")
//	//	if err != nil {
//	//		fmt.Println(err)
//	//	} else {
//	//		uid, _ = strconv.ParseInt(ck_uid, 10, 64)
//	//	}
//	str_goodsid := this.Input().Get("goodsid")
//	goodsid, _ := strconv.ParseInt(str_goodsid, 10, 64)
//	count, _ := strconv.ParseInt(this.Input().Get("count"), 10, 32)
//	if goodsid == 0 || count == 0 {
//		fmt.Println(goodsid, count)
//		return
//	}
//	v := time.Now().Unix()
//	nowDate := strconv.FormatInt(v, 10)
//	//	spcar := &models.Shopcar{Uid: uid, Goodsid: goodsid, Count: count, CreateDate: nowDate}
//	spcar := &models.Shopcar{Uid: 1, Goodsid: goodsid, Count: int32(count), CreateDate: nowDate}
//	count, err := models.OpShopcar(spcar, "add")
//	if err != nil {
//		fmt.Println(err)
//	}

//	var str string = "/buypage?goodsid=" + str_goodsid
//	this.Redirect(str, 303)
//}

func (this *BuyPageController) GetPageData(goodsid int64) {
	//	var uid int64 = 0
	//	ck_uid, err := this.Ctx.Request.Cookie("uid")
	//	if err != nil {
	//		fmt.Println(err)
	//	} else {
	//		uid, _ = strconv.ParseInt(ck_uid, 10, 64)
	//	}
	buypictures, _, err := models.GetPicturesByGoodsid(goodsid, "buypage", "slider", 1)
	if err != nil {
		fmt.Println(err)
		this.Data["CarCount"] = 0
	}

	goodsinfo, err1 := models.GetGoodsInfo(goodsid)
	if err1 != nil {
		fmt.Println(err)
	}
	//models.GetShopcarCount(uid)
	count := models.GetShopcarCount(1)
	this.Data["CarCount"] = count
	this.Data["GoodsInfo"] = goodsinfo
	this.Data["SliderPs"] = buypictures

}
