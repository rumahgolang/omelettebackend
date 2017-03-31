package models

type Category struct {
	Id         int    `json:"id"`
	MerchantId int    `json:"merchantId" form:"merchantId"`
	Name       string `json:"name" form:"name"`
}
