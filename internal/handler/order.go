package handler

import (
	"ServiceT/internal/model"
	"ServiceT/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	orderService *service.OrderService
}

func NewOrderHandler(os *service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: os}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order model.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.orderService.CreateOrder(&order); err != nil {
		http.Error(w, "Could not create order", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.orderService.GetAllOrders()
	if err != nil {
		http.Error(w, "Could not retrieve orders", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/orders/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	order, err := h.orderService.GetOrderByID(id)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetOrdersByUserID(w http.ResponseWriter, r *http.Request) {
	userIDStr := mux.Vars(r)["user_id"] // Извлекаем user_id из URL
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}
	orders, err := h.orderService.GetOrdersByUserID(userID) // Вызов сервиса
	if err != nil {
		http.Error(w, "Could not retrieve orders", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orders)
}
