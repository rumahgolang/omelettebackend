package models

type Orders struct {
	Id               int64  `db:"id" json:"id"`
	ItemId           int64  `db:"item_id" json:"item_id"`
	ItemName         string `db:"item_name" json:"item_name"`
	ItemPrice        int64  `db:"item_price" json:"item_price"`
	MerchantId       int64  `db:"merchant_id" json:"merchant_id"`
	MerchantName     string `db:"merchant_name" json:"merchant_name"`
	MemberId         string `db:"member_id" json:"member_id"`
	Qty              int8   `db:"qty" json:"qty"`
	ConfirmationCode string `db:"confirmation_code" json:"confirmation_code"`
}
