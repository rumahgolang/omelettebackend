package routers

import (
	"github.com/astaxie/beego"
	"github.com/fajarpnugroho/omelettebackend/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/dashboard", &controllers.DashboarController{})
	beego.Router("/merchant", &controllers.MerchantController{})
	beego.Router("/maps", &controllers.MapsController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/menu", &controllers.MenuController{})
}
