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

func GetPicturesByGoodsid(Goodsid int64, Where string, Type string, IsShow int8) ([]*Picture, int64, error) {
	ptc := make([]*Picture, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("picture")
	n, err := qs.Filter("goodsid", Goodsid).Filter("type", Type).Filter("isshow", IsShow).Filter("where", Where).All(&ptc)
	if err != nil {
		fmt.Println(err)
		return nil, 0, err
	}

	if n == 0 {
		return nil, n, err
	}
	return ptc, n, err
}

func GetGoodsInfo(goodsid int64) (*Goodsinfo, error) {
	var goods *Goodsinfo = &Goodsinfo{}
	o := orm.NewOrm()
	qs := o.QueryTable("goodsinfo")
	err := qs.Filter("goodsid", goodsid).One(goods)
	return goods, err
}

func OpShopcars(sc []*Shopcar, op string) (int64, error) {
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
				return 0, err
			}
		}
	}

	var count int64 = 0

	if len(sc) != 0 {
		count = GetShopcarCount(sc[0].Uid)
	}

	return count, nil
}

func OpShopcar(sc *Shopcar, op string) (int64, error) {
	o := orm.NewOrm()
	if op == "add" {
		dbShopcar, _ := GetShopcarByGoodsid(sc.Goodsid)
		if dbShopcar != nil {
			//			fmt.Println(sc)
			dbShopcar.Count += sc.Count
			o.Update(dbShopcar)
		} else {
			o.Insert(sc)
		}
	} else if op == "del" {
		dbShopcar, err := GetShopcarByGoodsid(sc.Goodsid)
		if dbShopcar != nil {
			o.Delete(dbShopcar)
		} else {
			return 0, err
		}
	}

	count := GetShopcarCount(sc.Uid)
	return count, nil
}

func GetShopcarCount(uid int64) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable("shopcar")
	count, err := qs.Filter("uid", uid).Count()
	if err != nil {
		return 0
	}
	return count
}

func GetShopcarByGoodsid(goodsid int64) (*Shopcar, error) {
	Spcar := &Shopcar{}
	o := orm.NewOrm()
	qs := o.QueryTable("shopcar")
	err := qs.Filter("goodsid", goodsid).One(Spcar)
	if err != nil {
		return nil, err
	}
	return Spcar, nil
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

		pictures, count, err2 := GetPicturesByGoodsid(v.Goodsid, "shopcar", "shopcarlist", 1)
		if err2 != nil {
			fmt.Println(err2)
			return nil
		}
		if count == 0 {
			fmt.Println("[func GetShopCarsData] Pictures is nil")
			return nil
		}
		CarDatas[k] = &ShopCarData{Goodsid: v.Goodsid, Name: goodsinfo.Name, Count: v.Count, Price: goodsinfo.Price,
			Createdate: v.CreateDate, Imgsrc: pictures[0].Imgsrc, Tourl: goodsinfo.Tourl}
	}

	return CarDatas
}

func GetCategorys() ([]*Category, int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	categorys := make([]*Category, 0)
	n, err := qs.All(&categorys)
	if err != nil {
		return nil, n, err
	}
	if n == 0 {
		return nil, n, err
	}

	return categorys, n, err
}

func GetGoodsinfoByCategory() (map[int32][]*ShopGoodsinfo, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Category")
	categorys := make([]*Category, 0)
	n, err := qs.All(&categorys)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, err
	}
	map_Shopinfo := make(map[int32][]*ShopGoodsinfo, n)
	for _, v := range categorys {
		goodsinfos := make([]*Goodsinfo, 0)
		qs = o.QueryTable("goodsinfo")
		count, _ := qs.Filter("category_name", v.Name).All(&goodsinfos)
		if count == 0 {
			continue
		}
		shopgoodsinfo := make([]*ShopGoodsinfo, count)

		for k, gds := range goodsinfos {
			pic, picCount, _ := GetPicturesByGoodsid(gds.Goodsid, "shoppage", "goodsshow", 1)
			tmpshopinfo := &ShopGoodsinfo{}
			if picCount != 0 {
				tmpshopinfo = &ShopGoodsinfo{Goodsid: gds.Goodsid, Price: gds.Price, Content: gds.Content, Name: gds.Name, Imgsrc: pic[0].Imgsrc, Tourl: pic[0].Tourl}
			} else {
				tmpshopinfo = &ShopGoodsinfo{Goodsid: gds.Goodsid, Price: gds.Price, Content: gds.Content, Name: gds.Name, Imgsrc: "#", Tourl: "#"}
			}
			shopgoodsinfo[k] = tmpshopinfo
		}
		map_Shopinfo[v.Categoryid] = shopgoodsinfo
	}

	return map_Shopinfo, err
}

func GetBooklist(uid int64) ([]*Book, int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("book")
	booklist := make([]*Book, 0)

	n, err := qs.Filter("uid", uid).OrderBy("statues").Limit(20).All(&booklist)
	if err != nil || n == 0 {
		return nil, n, err
	}

	return booklist, n, err
}

func CheckUserName(uname string) (bool, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("character")
	count, err := qs.Filter("uname", uname).Count()
	if err != nil {
		return false, err
	}

	if count != 0 {
		return false, err
	}
	return true, nil
}

func CreateUser(uname string, pw string, nickname string, email string, tel string, platform string) (int64, bool, error) {
	newUser := &Character{Uname: uname, Nickname: nickname, Pw: pw, Type: 0, Email: email, Phone: tel, Platform: platform, Token: pw}
	o := orm.NewOrm()
	qs := o.QueryTable("character")
	count, err := qs.Filter("uname", uname).Count()
	if err != nil {
		return -1, false, err
	}

	if count != 0 {
		return -1, false, err
	}

	uid, err1 := o.Insert(newUser)
	if err1 != nil {
		return -1, false, err1
	}
	return uid, true, err1
}

func UserLogin(uname string, pw string) *Character {
	o := orm.NewOrm()
	qs := o.QueryTable("character")
	count, err := qs.Filter("uname", uname).Count()

	if err != nil {
		return nil
	}
	if count == 0 {
		return nil
	}

	user := &Character{}
	err1 := qs.Filter("uname", uname).One(user)
	if err1 != nil {
		return nil
	}

	return user
}

func GetUserInfo(uname string, token string) (*Character, int32) {
	o := orm.NewOrm()
	qs := o.QueryTable("character")
	user := &Character{}
	err := qs.Filter("uname", uname).One(user)
	if err != nil {
		return nil, 1
	}

	if user.Token != token {
		return nil, 2
	}

	return user, 0
}

func GetOnlineUser(uname string, token string) (*Online, int32) {
	o := orm.NewOrm()
	qs := o.QueryTable("online")
	useronline := &Online{}
	err := qs.Filter("uname", uname).One(useronline)
	if err != nil {
		return nil, 1
	}
	if useronline.Token != token {
		return nil, 2
	}
	return useronline, 0
}

func OperatOnline(useronline *Online, op string) {
	o := orm.NewOrm()
	qs := o.QueryTable("online")
	count, _ := qs.Filter("uname", useronline.Uname).Count()

	switch op {
	case "add":
		{
			if count != 0 {
				o.Update(useronline)
			} else {
				o.Insert(useronline)
			}
		}
	case "del":
		{
			o.Delete(useronline)
		}

	case "update":
		{
			o.Update(useronline)
		}
	}
}

func GetAddress(uid int64) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("address")
	Addresses := make([]*Address, 0)
	count, err := qs.Filter("uid", uid).All(&Addresses)
	return count, err
}
