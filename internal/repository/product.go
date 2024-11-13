package repository

import (
	"ServiceT/internal/model"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *model.Product) error {
	_, err := r.db.NamedExec(`INSERT INTO products (name, price, quantity) VALUES (:name, :price, :quantity)`, product)
	return err
}

func (r *ProductRepository) GetAll() ([]model.Product, error) {
	var products []model.Product
	err := r.db.Select(&products, "SELECT * FROM products")
	return products, err
}

func (r *ProductRepository) GetByID(id int) (*model.Product, error) {
	var product model.Product
	err := r.db.Get(&product, "SELECT * FROM products WHERE id=$1", id)
	return &product, err
}

func (r *ProductRepository) Update(product *model.Product) error {
	_, err := r.db.NamedExec(`UPDATE products SET name=:name, price=:price, quantity=:quantity WHERE id=:id`, product)
	return err
}

func (r *ProductRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id=$1", id)
	return err
}
