package controllers

import (
	"FD_WebServer/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type ShopCarController struct {
	beego.Controller
}

func (this *ShopCarController) Get() {
	//	Uid_ck, err := this.Ctx.Request.Cookie("uid")
	haveCookie := false
	haveInfo := false
	var Uid int64 = 1
	//	if err != nil {
	//		fmt.Println("ShopCarController Get error:", err)
	//	} else {
	var err1 error
	//	Uid, err1 = strconv.ParseInt(Uid_ck.Value, 10, 64)
	if err1 != nil {
		fmt.Println("ShopCarController cookie error:", err1)
	} else {
		haveCookie = true
	}
	//	}

	ShopcarsData := models.GetShopCarsData(Uid)
	if ShopcarsData != nil {
		this.Data["ShopcarData"] = ShopcarsData
		haveInfo = true
	} else {
		this.Data["ShopcarData"] = "nil"
	}

	this.Data["haveCookie"] = haveCookie
	this.Data["haveInfo"] = haveInfo
	this.Data["ShopcarData"] = ShopcarsData
	this.Data["ServerHost"] = beego.AppConfig.String("ServerHost")
	this.Data["ServerPort"] = beego.AppConfig.String("ServerPort")
	this.TplName = "tab-subpage-shop-car.html"
}

func (this *ShopCarController) Remove() {
	options := this.Input().Get("op")
	strGoodsid := this.Input().Get("goodsid")
	goodsid, _ := strconv.ParseInt(strGoodsid, 10, 64)
	fmt.Println(options)
	fmt.Println(goodsid)
	switch options {
	case "del":
		spcar := &models.Shopcar{Uid: 1, Goodsid: goodsid}
		fmt.Println(spcar)
		_, err := models.OpShopcar(spcar, "del")
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("Default")
	}
	this.TplName = "tab-subpage-shop-car.html"
}
