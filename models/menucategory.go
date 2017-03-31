package models

// MenuCategory struct of menu category
type MenuCategory struct {
	ID         int64  `json:"id"`
	MerchantID int64  `json:"merchant_id"`
	Name       string `json:"name"`
}
