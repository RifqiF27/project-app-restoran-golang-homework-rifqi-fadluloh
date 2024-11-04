package model

type Discount struct {
	ID         uint16  `json:"id"`
	Code       string  `json:"code"`
	Percentage float64 `json:"percentage"`
}
