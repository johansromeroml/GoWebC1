package service

import "class1api/internal/repository"

type ProductService struct {
	pr repository.ProductJSONRepo
}

func NewProductService(repo repository.ProductJSONRepo) *ProductService {
	return &ProductService{pr: repo}
}
