package controllers

import (
	"github.com/astaxie/beego"
)

type MerchantController struct {
	beego.Controller
}

func (c *MerchantController) Get() {
	c.Data["merchant"] = "active"
	c.Data["titlePage"] = "Merchant"
	c.Layout = "main.tpl"
	c.TplName = "merchant.tpl"
}
