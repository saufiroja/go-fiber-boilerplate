package auth

import (
	"project/go-fiber-boilerplate/interfaces"
	"project/go-fiber-boilerplate/models/dto"
	"project/go-fiber-boilerplate/utils"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	repo     interfaces.AuthRepository
	validate *validator.Validate
}

func NewAuthService(repo interfaces.AuthRepository) interfaces.AuthService {
	return &Service{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *Service) Register(user *dto.Register) error {
	// validation
	err := s.validate.Struct(user)
	if err != nil {
		return utils.HandlerError(err)
	}

	// hash password
	hash := utils.HashPassword(user.Password)
	user.Password = hash

	return s.repo.Register(user)
}

func (s *Service) Login(user *dto.Login) (*dto.LoginResponse, error) {
	// validation
	err := s.validate.Struct(user)
	if err != nil {
		return nil, utils.HandlerError(err)
	}

	// check email
	res, err := s.repo.Login(user.Email)
	if err != nil {
		return nil, utils.HandlerErrorCustom(404, "email not found")
	}

	err = utils.ComparePassword(res.Password, user.Password)
	if err != nil {
		return nil, utils.HandlerErrorCustom(400, "wrong password")
	}

	accessToken, expiredAccessToken, err := utils.GenerateAccessToken(res.ID, res.Email, res.FullName)
	if err != nil {
		return nil, utils.HandlerErrorCustom(500, "failed to generate access token")
	}

	refreshToken, expiredRefreshToken, err := utils.GenerateRefreshToken(res.ID, res.Email, res.FullName)
	if err != nil {
		return nil, utils.HandlerErrorCustom(500, "failed to generate refresh token")
	}

	token := &dto.LoginResponse{
		AccessToken:         accessToken,
		AccessTokenExpired:  expiredAccessToken,
		RefreshToken:        refreshToken,
		RefreshTokenExpired: expiredRefreshToken,
	}

	return token, nil
}
