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
	DBName := beego.AppConfig.String("MySqlDBName")
	DBsrc := user + ":" + DBpw + "@(" + host + ":" + port + ")/" + DBName + "?charset=utf8"
	orm.RegisterModel(new(Character), new(Goodsinfo), new(Book), new(Picture),
		new(EverydaySales), new(Evaluate), new(Poster), new(Rechage),
		new(Collect), new(DiscountCoupon), new(ExpenseLog), new(Online), new(Address))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", DBsrc)
	orm.RunSyncdb("default", false, true)
}
