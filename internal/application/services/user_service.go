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
func (s *UserService) SaveUser(name string, age int) error {
	return s.repo.SaveUser(name, age)
}

// GetUsers retorna a lista de usuários cadastrados
func (s *UserService) GetUsers() ([]domain.User, error) {
	users, err := s.repo.GetAllUsers() // Este método deve ser implementado no seu repositório
	if err != nil {
		return nil, err
	}
	return users, nil
}
