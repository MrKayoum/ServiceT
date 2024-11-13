package repository

import (
	"ServiceT/internal/model"
	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository { // Обратите внимание на *OrderRepository
	return &OrderRepository{db: db} // Здесь всё правильно
}

// Создание нового заказа
func (r *OrderRepository) Create(order *model.Order) error {
	_, err := r.db.NamedExec(`INSERT INTO orders (user_id, product_id, quantity, status) VALUES (:user_id, :product_id, :quantity, :status)`, order)
	return err
}

// Получение всех заказов
func (r *OrderRepository) GetAll() ([]model.Order, error) {
	var orders []model.Order
	err := r.db.Select(&orders, "SELECT * FROM orders")
	return orders, err
}

// Новый метод - Получение заказов по user_id
func (r *OrderRepository) GetByUserID(userID int) ([]model.Order, error) { // Новый метод
	var orders []model.Order
	err := r.db.Select(&orders, "SELECT * FROM orders WHERE user_id=$1", userID)
	return orders, err
}

// Получение заказа по ID
func (r *OrderRepository) GetByID(orderID int) (*model.Order, error) {
	var order model.Order
	err := r.db.Get(&order, "SELECT * FROM orders WHERE id=$1", orderID)
	return &order, err
}
