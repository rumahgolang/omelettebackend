package controllers

import (
	"github.com/astaxie/beego"
)

type DashboarController struct {
	beego.Controller
}

func (c *DashboarController) Get() {
	c.Data["dashboard"] = "active"
	c.Data["titlePage"] = "Dashboard"
	c.Layout = "main.tpl"
	c.TplName = "dashboard.tpl"
}
