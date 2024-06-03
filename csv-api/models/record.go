package models

type Record struct {
	Item     string  `json:"item"`
	Value    string  `json:"value"`
	Category string  `json:"category"`
	Date     string  `json:"date"`
	ID       int     `json:"id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
