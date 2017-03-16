package controllers

import (
	"github.com/astaxie/beego"
)

type MapsController struct {
	beego.Controller
}

func (c *MapsController) Get() {
	c.Data["maps"] = "active"
	c.Data["titlePage"] = "Maps"
	c.Layout = "main.tpl"
	c.TplName = "maps.tpl"
}
