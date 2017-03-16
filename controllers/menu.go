package controllers

import (
	"github.com/astaxie/beego"
)

type MenuController struct {
	beego.Controller
}

func (c *MenuController) Get() {
	c.Data["menu"] = "active"
	c.Data["titlePage"] = "Menu"
	c.Layout = "main.tpl"
	c.TplName = "menu.tpl"
}
