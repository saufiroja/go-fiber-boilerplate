package auth

import (
	"project/go-fiber-boilerplate/dto"
	"project/go-fiber-boilerplate/interfaces"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	repo     interfaces.UserRepository
	validate *validator.Validate
}

func NewAuthService(repo interfaces.UserRepository) interfaces.UserService {
	return &Service{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *Service) Register(user *dto.Register) error {
	return s.repo.Register(user)
}

// func (s *Service) Login(email, password string) error {
// first query database
// check if tidak ada  return error
// check if password salah return error
// data dari query database
// access token sama refresh
// return login response
// }
