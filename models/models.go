package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func RegistDB() {
	host := beego.AppConfig.String("MySqlHost")
	port := beego.AppConfig.String("MySqlPort")
	user := beego.AppConfig.String("MySqlUser")
	DBpw := beego.AppConfig.String("MySqlPassWord")
	if DBpw == "123456" {
		DBpw = "arvin"
	} else {
		DBpw = "123456"
	}
	DBName := beego.AppConfig.String("MySqlDBName")
	DBsrc := user + ":" + DBpw + "@(" + host + ":" + port + ")/" + DBName + "?charset=utf8"
	orm.RegisterModel(new(Character), new(Goodsinfo), new(Book), new(Picture), new(EverydaySales),
		new(Evaluate), new(Rechage), new(Collect), new(DiscountCoupon), new(ExpenseLog), new(Online),
		new(Address), new(Poster), new(Postergift), new(GiftLog), new(Shopcar), new(Category))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", DBsrc)
	orm.RunSyncdb("default", false, true)
}
