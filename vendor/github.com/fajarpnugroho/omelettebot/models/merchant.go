package models

type Merchant struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	Rating        float64 `json:"rating"`
	Location      string  `json:"location"`
	RangePrice    string  `json:"range_price"`
	OpenCloseInfo string  `json:"open_close_info"`
	Pict          string  `json:"pict"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Distance      float64 `json:"distance"`
}
