package handler

import (
	"ServiceT/internal/model"
	"ServiceT/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(ps *service.ProductService) *ProductHandler {
	return &ProductHandler{productService: ps}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.productService.CreateProduct(&product); err != nil {
		http.Error(w, "Could not create product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		http.Error(w, "Could not retrieve products", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	product, err := h.productService.GetProductByID(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	product.ID = id // Устанавливаем ID для обновляемого продукта
	if err := h.productService.UpdateProduct(&product); err != nil {
		http.Error(w, "Could not update product", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.productService.DeleteProduct(id); err != nil {
		http.Error(w, "Could not delete product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent) // Ответ без содержимого
}
