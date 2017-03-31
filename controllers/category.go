package controllers

import (
	"log"
	"strconv"

	"github.com/fajarpnugroho/omelettebackend/models"

	"github.com/astaxie/beego"
	"github.com/fajarpnugroho/omelettebackend/services"
	"github.com/fajarpnugroho/omelettebackend/utils"
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

func (c *CategoryController) ShowCategoryForm() {
	log.Println("Add Category")
	c.Data["category"] = "active"
	c.Data["titlePage"] = "Category Menu"
	c.Layout = "main.tpl"
	c.TplName = "category.tpl"

	merchantID := c.Ctx.Input.Param(":merchantid")
	c.Data["merchantId"] = merchantID
}

// Add add new category
func (c *CategoryController) Add() {
	log.Println("Add Category")
	c.Data["category"] = "active"
	c.Data["titlePage"] = "Category Menu"
	c.Layout = "main.tpl"
	c.TplName = "category.tpl"

	// cat := models.MenuCategory{}
	// if err := c.ParseForm(&cat); err != nil {
	// 	//handle error
	// 	c.Ctx.ResponseWriter.Write(utils.STR_ERROR_REQUEST_INVALID)
	// 	return
	// }

	categoryName := c.GetString("category_name", "")
	merchantID, _ := strconv.ParseInt(c.Ctx.Input.Param(":merchantid"), 10, 64)

	cat := models.MenuCategory{
		MerchantID: merchantID,
		Name:       categoryName,
	}

	errResp := services.AddedNewCategory(&cat)

	if errResp != nil {
		c.Ctx.ResponseWriter.Write(utils.STR_CATEGORY_RESPONSE_ERROR_ADD_CATEGORY)
		return
	}

	c.Redirect("/category/"+c.Ctx.Input.Param(":merchantid")+"?saved=true", 302)
}

// GetAllMenuByMerchantId get all category menu by merchant_id
func (c *CategoryController) GetAllMenuByMerchantId() {
	log.Println("Get All Category Menu By Merchant")
	c.Data["category"] = "active"
	c.Layout = "main.tpl"
	c.TplName = "category.tpl"

	merchantID := c.Ctx.Input.Param(":merchantid")

	listCategory, errQuery := services.GetAllMenuByMerchantId(merchantID)

	if errQuery != nil {
		c.Ctx.ResponseWriter.Write(utils.STR_CATEGORY_RESPONSE_ERROR_GET_CATEGORIES)
		return
	}

	merchant, _ := services.GetDetailMerchant(merchantID)
	c.Data["titlePage"] = merchant.Name
	c.Data["merchant"] = merchant
	c.Data["listCategories"] = listCategory
}

func (c *CategoryController) UpdateCategoryByMerchantId() {
	log.Println("Get All Menu By Merchant")
	c.Data["category"] = "active"
	c.Data["titlePage"] = "Category Menu"
	c.Layout = "main.tpl"
	c.TplName = "category.tpl"

	merchantID := c.Ctx.Input.Param(":merchantid")
	categoryID := c.Ctx.Input.Param(":categoryid")

	cat := models.MenuCategory{}
	cat.ID, _ = strconv.ParseInt(categoryID, 10, 64)
	cat.MerchantID, _ = strconv.ParseInt(merchantID, 10, 64)

	if err := c.ParseForm(&cat); err != nil {
		//handle error
		c.Ctx.ResponseWriter.Write(utils.STR_ERROR_REQUEST_INVALID)
		return
	}

	listCategory, errQuery := services.UpdateCategoryByMerchantId(&cat)

	if errQuery != nil {
		c.Ctx.ResponseWriter.Write(utils.STR_CATEGORY_RESPONSE_ERROR_GET_CATEGORIES)
		return
	}
	c.Data["listCategories"] = listCategory
}

func (c *CategoryController) DeleteCategoryByMerchantId() {
	log.Println("Get All Menu By Merchant")
	c.Data["category"] = "active"
	c.Data["titlePage"] = "Category Menu"
	c.Layout = "main.tpl"
	c.TplName = "category.tpl"

	merchantID := c.Ctx.Input.Param(":merchantid")
	categoryID := c.Ctx.Input.Param(":categoryid")

	cat := models.MenuCategory{}
	cat.ID, _ = strconv.ParseInt(categoryID, 10, 64)
	cat.MerchantID, _ = strconv.ParseInt(merchantID, 10, 64)

	errQuery := services.DeleteCategoryByMerchantId(&cat)

	if errQuery != nil {
		c.Ctx.ResponseWriter.Write(utils.STR_CATEGORY_RESPONSE_ERROR_GET_CATEGORIES)
		return
	}

	c.Ctx.ResponseWriter.Write(utils.STR_RESPONSE_OK)
}
