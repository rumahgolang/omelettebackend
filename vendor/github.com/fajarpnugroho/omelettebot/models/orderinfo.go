package models

type OrderInfo struct {
	MerchantId   int64  `db:"merchant_id" json:"merchant_id"`
	MerchantName string `db:"merchant_name" json:"merchant_name"`
	MemberId     string `db:"member_id" json:"member_id"`
	LineId       string `db:"line_id" json:"line_id"`
	DisplayName  string `db:"display_name" json:"display_name"`
	OrderDate    int64  `db:"order_date" json:"order_date"`
	ModifiedDate int64  `db:"modified_date" json:"modified_date"`
}
