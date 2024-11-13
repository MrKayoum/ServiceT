package model

type Product struct {
	ID       int     `db:"id" json:"id"`
	Name     string  `db:"name" json:"name"`
	Price    float64 `db:"price" json:"price"`
	Quantity int     `db:"quantity" json:"quantity"`
}
