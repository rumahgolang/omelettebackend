package models

type Member struct {
	Id          int64  `db:"id" json:"id"`
	MemberId    string `db:"member_id" json:"member_id"`
	LineId      string `db:"line_id" json:"line_id"`
	DisplayName string `db:"display_name" json:"display_name"`
}
