package usecases

import "bronze/internal/domain"

type UserRepository interface {
	SaveUser(name string, age int) error
	GetAllUsers() ([]domain.User, error)
}
