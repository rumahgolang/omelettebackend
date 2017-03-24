package routers

import (
	"github.com/astaxie/beego"
	"github.com/fajarpnugroho/omelettebackend/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/dashboard", &controllers.DashboarController{})
	beego.Router("/merchant", &controllers.MerchantController{})
	beego.Router("/merchant/:id/add", &controllers.MerchantController{}, "get:AddMerchant")
	beego.Router("/merchant/:id/save", &controllers.MerchantController{}, "post:SaveMerchant")
	beego.Router("/places", &controllers.PlacesController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/menu", &controllers.MenuController{})
	beego.Router("/result", &controllers.ResultController{})
}
