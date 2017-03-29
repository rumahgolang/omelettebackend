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

func (c *CategoryController) Add() {
	log.Println("Add Category")
	c.Data["category"] = "active"
	c.Data["titlePage"] = "Category Menu"
	c.Layout = "main.tpl"
	c.TplName = "category.tpl"

	cat := models.Category{}
	if err := c.ParseForm(&cat); err != nil {
		//handle error
		c.Ctx.ResponseWriter.Write(utils.STR_ERROR_REQUEST_INVALID)
		return
	}

	errResp := services.AddedNewCategory(&cat)

	if errResp != nil {
		c.Ctx.ResponseWriter.Write(utils.STR_CATEGORY_RESPONSE_ERROR_ADD_CATEGORY)
		return
	}
	c.Ctx.ResponseWriter.Write(utils.STR_RESPONSE_OK)
}

func (c *CategoryController) GetAllMenuByMerchantId() {
	log.Println("Get All Menu By Merchant")
	c.Data["category"] = "active"
	c.Data["titlePage"] = "Category Menu"
	c.Layout = "main.tpl"
	c.TplName = "category.tpl"

	merchantID := c.Ctx.Input.Param(":merchantid")

	listCategory, errQuery := services.GetAllMenuByMerchantId(merchantID)

	if errQuery != nil {
		c.Ctx.ResponseWriter.Write(utils.STR_CATEGORY_RESPONSE_ERROR_GET_CATEGORIES)
		return
	}
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

	cat := models.Category{}
	cat.Id, _ = strconv.Atoi(categoryID)
	cat.MerchantId, _ = strconv.Atoi(merchantID)

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

	cat := models.Category{}
	cat.Id, _ = strconv.Atoi(categoryID)
	cat.MerchantId, _ = strconv.Atoi(merchantID)

	errQuery := services.DeleteCategoryByMerchantId(&cat)

	if errQuery != nil {
		c.Ctx.ResponseWriter.Write(utils.STR_CATEGORY_RESPONSE_ERROR_GET_CATEGORIES)
		return
	}

	c.Ctx.ResponseWriter.Write(utils.STR_RESPONSE_OK)
}
