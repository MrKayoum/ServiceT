package model

type Order struct {
	ID        int    `db:"id" json:"id"`
	UserID    int    `db:"user_id" json:"user_id"`
	ProductID int    `db:"product_id" json:"product_id"`
	Quantity  int    `db:"quantity" json:"quantity"`
	Status    string `db:"status" json:"status"`
}
