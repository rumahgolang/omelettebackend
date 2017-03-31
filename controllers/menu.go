package controllers

import (
	"github.com/astaxie/beego"
)

// MenuController menu controllers
type MenuController struct {
	beego.Controller
}

// Get index page
func (c *MenuController) Get() {
	c.Data["menu"] = "active"
	c.Data["titlePage"] = "Menu"
	c.Layout = "main.tpl"
	c.TplName = "menu.tpl"
}
