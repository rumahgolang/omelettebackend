package models

type MenuCategory struct {
	Id         int64  `json:"id"`
	MerchantId int64  `json:"merchant_id"`
	Name       string `json:"name"`
}
