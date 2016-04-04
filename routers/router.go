package routers

import (
	"FD_WebServer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/home", &controllers.MainController{})
	beego.Router("/shop", &controllers.ShopController{})
	beego.Router("/shop-car", &controllers.ShopCarController{})
	beego.Router("/book", &controllers.BookController{})
}
