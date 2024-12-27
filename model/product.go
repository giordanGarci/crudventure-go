package model

type Product struct {
	ID    int64   `json:"id_product"`
	Name  string  `json:"product_name"`
	Price float64 `json:"price"`
}
