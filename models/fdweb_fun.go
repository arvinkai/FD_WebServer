// fdweb_fun
package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

func GetPicture(Where string, Type string, IsShow int8) ([]*Picture, int64, error) {
	var ptc []*Picture = make([]*Picture, 0)

	o := orm.NewOrm()
	qs := o.QueryTable("picture")
	n, err := qs.Filter("type", Type).Filter("isshow", IsShow).Filter("where", Where).All(&ptc)
	if err != nil {
		fmt.Println(err)
		return nil, n, err
	}
	if n == 0 {
		return nil, n, err
	}

	return ptc, n, err
}

func GetGoodsInfo(goodsid int64) (*Goodsinfo, error) {
	var goods *Goodsinfo = &Goodsinfo{} //new(Goodsinfo)
	o := orm.NewOrm()
	qs := o.QueryTable("goodsinfo")
	err := qs.Filter("goodsid", goodsid).One(goods)
	return goods, err
}

func GetShopcarGoodsids(uid int64) ([]*Cargoods, int64, error) {
	Cargoods := make([]*Goodsinfo, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("shopcar")
	n, err := qs.Filter("uid", uid).All(&Cargoods)
	if err != nil {
		return nil, n, err
	}

	if n == 0 {
		return nil, 0, err
	}

	return Cargoods, n, err
}

func GetShopCar(uid int64) []*ShopCarData {
	CarData := make([]*ShopCarData, 0)
	Cargoodses, n, err := GetShopcarGoodsids(uid)

}
