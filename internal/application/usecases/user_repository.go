package usecases

import "bronze/internal/domain"

type UserRepository interface {
	SaveProduct(name string, age float64) error
	GetAllProducts() ([]domain.Product, error)
}
