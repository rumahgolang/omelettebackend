package models

type MenuItem struct {
	Id             int64  `json:"menu_id"`
	MerchantId     int64  `json:"merchant_id"`
	MenuCategoryId int64  `json:"menu_category_id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Price          int64  `json:"price"`
	Pict           string `json:"pict"`
}
