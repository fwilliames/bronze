package usecases

import "bronze/internal/domain"

type UserRepository interface {
	SaveProduct(name, data, market string, value, totalValue float64, quantity int64) error
	SaveMarket(name string) error
	SaveData(data string) error
	GetAllProducts() ([]domain.Product, error)
	GetAllProductsbyFilter(filter string) ([]domain.Product, error)
	GetUniqueDates() ([]string, error)
	GetAllMarkets() ([]string, error)
	GetAllDates() ([]domain.Data, error)
}
