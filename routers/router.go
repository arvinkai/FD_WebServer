package routers

import (
	"FD_WebServer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.HomePageController{})
	beego.Router("/shop", &controllers.ShopController{})
	beego.Router("/shopcar", &controllers.ShopCarController{})
	beego.AutoRouter(&controllers.ShopCarController{})
	beego.Router("/book", &controllers.BookController{})
	beego.Router("/setting", &controllers.UserSettingController{})
	beego.AutoRouter(&controllers.UserSettingController{})
	beego.Router("/buypage", &controllers.BuyPageController{})
	beego.AutoRouter(&controllers.BuyPageController{})
	beego.Router("/buyjump", &controllers.BuyJumpController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.AutoRouter(&controllers.LoginController{})
	beego.Router("/regist", &controllers.RegistController{})
	beego.AutoRouter(&controllers.RegistController{})
}
