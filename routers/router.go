package routers

import (
	"FD_WebServer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.HomeController{})
	beego.Router("/shop", &controllers.ShopController{})
	beego.Router("/order", &controllers.OrderController{})
}
