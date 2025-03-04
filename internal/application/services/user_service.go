package services

import (
	"bronze/internal/application/usecases"
	"bronze/internal/domain"
)

// UserService implementa os casos de uso
type UserService struct {
	repo usecases.UserRepository
}

// NewUserService cria uma nova instância do serviço
func NewUserService(repo usecases.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// SaveUser salva um usuário no banco
func (s *UserService) SaveUser(name string, value float64) error {
	return s.repo.SaveProduct(name, value)
}

// GetUsers retorna a lista de usuários cadastrados
func (s *UserService) GetUsers() ([]domain.Product, error) {
	users, err := s.repo.GetAllProducts() // Este método deve ser implementado no seu repositório
	if err != nil {
		return nil, err
	}
	return users, nil
}
