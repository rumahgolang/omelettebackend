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
	beego.Router("/category/", &controllers.CategoryController{})
	beego.Router("/category/:merchantid/show", &controllers.CategoryController{}, "get:ShowCategoryForm")
	beego.Router("/category/:merchantid/add", &controllers.CategoryController{}, "post:Add")
	beego.Router("/category/:merchantid", &controllers.CategoryController{}, "get:GetAllMenuByMerchantId")
	beego.Router("/category/:merchantid/:categoryid", &controllers.CategoryController{}, "put:UpdateCategoryByMerchantId")
	beego.Router("/category/:merchantid/:categoryid", &controllers.CategoryController{}, "delete:DeleteCategoryByMerchantId")
	beego.Router("/menu", &controllers.MenuController{})
	beego.Router("/result", &controllers.ResultController{})
}
