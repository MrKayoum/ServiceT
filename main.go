package main

import (
	"log"
	"net/http"

	"ServiceT/internal/handler"
	"ServiceT/internal/repository"
	"ServiceT/internal/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// Инициализация соединения с базой данных
	dsn := "user=postgres password=Egor73tagir dbname=user_service_db sslmode=disable"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)
	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	// Настройка маршрутов
	r := mux.NewRouter()
	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/login", userHandler.LoginUser).Methods("POST")
	r.HandleFunc("/user/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/products", productHandler.GetAllProducts).Methods("GET")
	r.HandleFunc("/products/{id}", productHandler.GetProductByID).Methods("GET")
	r.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	r.HandleFunc("/orders", orderHandler.GetAllOrders).Methods("GET")
	r.HandleFunc("/orders/{id}", orderHandler.GetOrderByID).Methods("GET")
	r.HandleFunc("/orders/user/{user_id}", orderHandler.GetOrdersByUserID).Methods("GET")
	log.Println("Server starting on :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
