package auth

import (
	"errors"
	"project/go-fiber-boilerplate/dto"
	"project/go-fiber-boilerplate/interfaces"
	"project/go-fiber-boilerplate/utils"

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
	hash := utils.HashPassword(user.Password)
	user.Password = hash
	return s.repo.Register(user)
}

func (s *Service) Login(email, password string) (*dto.LoginResponse, error) {
	user, err := s.repo.Login(email)
	if err != nil {
		return nil, errors.New("email or password is wrong")
	}

	_ = utils.ComparePassword(user.Password, password)

	accessToken, err := utils.GenerateAccessToken(user.ID, user.Email, user.FullName)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, user.Email, user.FullName)
	if err != nil {
		return nil, err
	}

	token := &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return token, nil
}
