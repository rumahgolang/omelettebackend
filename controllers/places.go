package controllers

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/fajarpnugroho/omelettebackend/models"
	"github.com/fajarpnugroho/omelettebackend/services"
)

type PlacesController struct {
	beego.Controller
}

// Get Home page
func (c *PlacesController) Get() {
	c.Data["maps"] = "active"
	c.Data["titlePage"] = "Places"
	c.Layout = "main.tpl"
	c.TplName = "places.tpl"
}

// Post Post home page
func (c *PlacesController) Post() {
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
	c.Data["titlePage"] = "Places"
	c.Layout = "main.tpl"
	c.TplName = "places.tpl"
}
