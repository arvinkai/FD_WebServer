package routers

import (
	"FD_WebServer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.HomePageController{})
	beego.Router("/shop", &controllers.ShopController{})
	beego.Router("/shop_car", &controllers.ShopCarController{})
	beego.AutoRouter(&controllers.ShopCarController{})
	beego.Router("/book", &controllers.BookController{})
	beego.Router("/setting", &controllers.UserSettingController{})
	beego.Router("/buypage", &controllers.BuyPageController{})
}
