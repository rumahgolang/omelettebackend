package controllers

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/fajarpnugroho/omelettebackend/models"
	"github.com/fajarpnugroho/omelettebackend/services"
)

type ResultController struct {
	beego.Controller
}

func (c *ResultController) Post() {
	data := models.Search{}
	if err := c.ParseForm(&data); err != nil {
		//handle error
		log.Println("Error parse form")
	}

	var zomatoResponse models.ZomatoRestaurantResponse
	services.SearchRestaurantByKeyword(&zomatoResponse, data.Keyword)

	for _, item := range zomatoResponse.Restaurants {
		log.Printf("%s", item.Restaurant.Name)
	}

	c.Data["restaurant"] = zomatoResponse.Restaurants

	c.Data["keyword"] = data.Keyword
	c.Data["maps"] = "active"
	c.Data["titlePage"] = "Places"
	c.Layout = "main.tpl"
	c.TplName = "places.tpl"
}
