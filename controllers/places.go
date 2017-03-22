package controllers

import "github.com/astaxie/beego"

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
