package main

import (
	"FD_WebServer/models"

	_ "FD_WebServer/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegistDB()
}

func main() {
	orm.Debug = true
	beego.Run()
}
