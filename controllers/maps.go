package controllers

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/fajarpnugroho/omelettebackend/models"
	"github.com/fajarpnugroho/omelettebackend/services"
)

type MapsController struct {
	beego.Controller
}

// Get Home page
func (c *MapsController) Get() {
	c.Data["maps"] = "active"
	c.Data["titlePage"] = "Maps"
	c.Layout = "main.tpl"
	c.TplName = "maps.tpl"
}

// Post Post home page
func (c *MapsController) Post() {
	data := models.Search{}
	if err := c.ParseForm(&data); err != nil {
		//handle error
		log.Println("Error parse form")
	}

	var zomatoResponse models.ZomatoRestaurantResponse
	services.SearchRestaurantByKeyword(&zomatoResponse, data.Keyword)

	c.Data["restaurant"] = zomatoResponse.Restaurants

	c.Data["keyword"] = data.Keyword
	c.Data["maps"] = "active"
	c.Data["titlePage"] = "Maps"
	c.Layout = "main.tpl"
	c.TplName = "maps.tpl"
}
