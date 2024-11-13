package service

import (
	"ServiceT/internal/model"
	"ServiceT/internal/repository"
)

type OrderService struct {
	orderRepo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{orderRepo: repo}
}

func (s *OrderService) CreateOrder(order *model.Order) error {
	return s.orderRepo.Create(order)
}

func (s *OrderService) GetAllOrders() ([]model.Order, error) {
	return s.orderRepo.GetAll()
}

func (s *OrderService) GetOrderByID(id int) (*model.Order, error) {
	return s.orderRepo.GetByID(id)
}

func (s *OrderService) GetOrdersByUserID(userID int) ([]model.Order, error) { // Функция возвращает срез заказов и ошибку
	orders, err := s.orderRepo.GetByUserID(userID)
	if err != nil {
		return nil, err // Возвращаем nil и ошибку
	}
	return orders, nil // Возвращаем найденные заказы и nil в качестве ошибки
}
