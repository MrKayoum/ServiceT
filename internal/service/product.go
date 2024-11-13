package service

import (
	"ServiceT/internal/model"
	"ServiceT/internal/repository"
)

type ProductService struct {
	productRepo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{productRepo: repo}
}

func (s *ProductService) CreateProduct(product *model.Product) error {
	return s.productRepo.Create(product)
}

func (s *ProductService) GetAllProducts() ([]model.Product, error) {
	return s.productRepo.GetAll()
}

func (s *ProductService) GetProductByID(id int) (*model.Product, error) {
	return s.productRepo.GetByID(id)
}

func (s *ProductService) UpdateProduct(product *model.Product) error {
	return s.productRepo.Update(product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.productRepo.Delete(id)
}
