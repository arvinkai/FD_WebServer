// fdweb_fun
package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

func GetPicturesByUse(Where string, Type string, IsShow int8) ([]*Picture, int64, error) {
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

func GetPictureByGoodsid(Goodsid int64, Where string, Type string, IsShow int8) (*Picture, error) {
	ptc := &Picture{}
	o := orm.NewOrm()
	qs := o.QueryTable("picture")
	err := qs.Filter("goodsid", Goodsid).Filter("type", Type).Filter("isshow", IsShow).Filter("where", Where).One(ptc)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return ptc, err
}

func GetGoodsInfo(goodsid int64) (*Goodsinfo, error) {
	var goods *Goodsinfo = &Goodsinfo{}
	o := orm.NewOrm()
	qs := o.QueryTable("goodsinfo")
	err := qs.Filter("goodsid", goodsid).One(goods)
	return goods, err
}

func OpShopcars(sc []*Shopcar, op string) error {
	o := orm.NewOrm()
	if op == "add" {
		for _, v := range sc {
			dbShopcar, _ := GetShopcarByGoodsid(v.Goodsid)
			if dbShopcar != nil {
				o.Update(v)
			} else {
				o.Insert(v)
			}
		}
	} else if op == "del" {
		for _, v := range sc {
			dbShopcar, err := GetShopcarByGoodsid(v.Goodsid)
			if dbShopcar != nil {
				o.Delete(v)
			} else {
				return err
			}
		}
	}

	return nil
}

func GetShopcarByGoodsid(goodsid int64) (*Shopcar, error) {
	Shopcar := &Shopcar{}
	o := orm.NewOrm()
	qs := o.QueryTable("shopcar")
	err := qs.Filter("goodsid", goodsid).One(&Shopcar)
	if err != nil {
		return nil, err
	}
	return Shopcar, nil
}

func GetShopcars(uid int64) ([]*Shopcar, int64, error) {
	Shopcars := make([]*Shopcar, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("shopcar")
	n, err := qs.Filter("uid", uid).All(&Shopcars)
	if err != nil {
		return nil, n, err
	}

	if n == 0 {
		return nil, 0, err
	}

	return Shopcars, n, err
}

func GetShopCarsData(uid int64) []*ShopCarData {
	Shopcars, n, err := GetShopcars(uid)
	if err != nil {
		fmt.Println("GetShopCar:", err)
		return nil
	}
	if n == 0 {
		fmt.Println("Fields cols:", n)
		return nil
	}
	CarDatas := make([]*ShopCarData, n)

	for k, v := range Shopcars {
		goodsinfo, err1 := GetGoodsInfo(v.Goodsid)
		if err1 != nil {
			fmt.Println(err1)
			return nil
		}

		picture, err2 := GetPictureByGoodsid(v.Goodsid, "shopcar", "shopcarlist", 1)
		if err2 != nil {
			fmt.Println(err2)
			return nil
		}

		CarDatas[k] = &ShopCarData{Goodsid: v.Goodsid, Name: goodsinfo.Name, Count: v.Count, Price: goodsinfo.Price,
			Createdate: v.CreateDate, Imgsrc: picture.Imgsrc, Tourl: goodsinfo.Tourl}
	}

	return CarDatas
}
