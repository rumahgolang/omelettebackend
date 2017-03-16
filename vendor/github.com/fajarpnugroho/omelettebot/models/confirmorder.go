package models

type ConfirmOrder struct {
	Id               int64  `db:"id" json:"id"`
	ConfirmationCode string `db:"confirmation_code" json:"confirmation_code"`
	Status           int    `db:"status" json:"status"`
	OrderDate        int64  `db:"order_date" json:"order_date"`
	ModifiedDate     int64  `db:"modified_date" json:"modified_date"`
}
