package controllers

import (
	"log"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/fajarpnugroho/omelettebackend/services"
	"github.com/fajarpnugroho/omelettebot/models"
)

type MerchantController struct {
	beego.Controller
}

// Get Index
func (c *MerchantController) Get() {
	c.Data["merchant"] = "active"
	c.Data["titlePage"] = "Merchant"
	c.Layout = "main.tpl"
	c.TplName = "merchant.tpl"
}

// AddMerchant add new merchant
func (c *MerchantController) AddMerchant() {
	merchantID := c.Ctx.Input.Param(":id")

	ID, _ := strconv.ParseInt(merchantID, 10, 64)

	merchantProfile := services.RestaurantDetail(ID)

	c.Data["profile"] = merchantProfile
	c.Data["merchant"] = "active"
	c.Data["titlePage"] = "Add Merchant"
	c.Layout = "main.tpl"
	c.TplName = "addmerchant.tpl"
}

// SaveMerchant save data mechant
func (c *MerchantController) SaveMerchant() {

	stringLat := c.GetString("latitude", "0")
	stringLong := c.GetString("longitude", "0")

	latitude, _ := strconv.ParseFloat(stringLat, 64)

	longitude, _ := strconv.ParseFloat(stringLong, 64)

	merchant := &models.Merchant{
		Name:          c.GetString("merchantname", ""),
		Location:      c.GetString("address", ""),
		RangePrice:    c.GetString("rangeprice", ""),
		OpenCloseInfo: c.GetString("opencloseinfo", ""),
		Latitude:      latitude,
		Longitude:     longitude,
	}

	log.Println(merchant)

	errAdd := services.AddedNewMerchant(merchant)

	if errAdd != nil {
		log.Println(errAdd)
	}

	c.Data["merchant"] = "active"
	c.Data["titlePage"] = "Data Saved"
	c.Layout = "main.tpl"
	c.TplName = "merchant.tpl"
}
