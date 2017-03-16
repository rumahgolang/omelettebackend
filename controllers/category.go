package controllers

import (
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	c.Data["category"] = "active"
	c.Data["titlePage"] = "Category Menu"
	c.Layout = "main.tpl"
	c.TplName = "category.tpl"
}
