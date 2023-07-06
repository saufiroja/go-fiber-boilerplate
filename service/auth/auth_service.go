package auth

import (
	"project/go-fiber-boilerplate/interfaces"
	"project/go-fiber-boilerplate/models/dto"
	"project/go-fiber-boilerplate/utils"

	"github.com/go-playground/validator/v10"
)

type authService struct {
	repoAuth interfaces.AuthRepository
	validate *validator.Validate
}

func NewAuthService(repoAuth interfaces.AuthRepository) interfaces.AuthService {
	return &authService{
		repoAuth: repoAuth,
		validate: validator.New(),
	}
}

func (s *authService) Register(user *dto.Register) error {
	// validation
	err := s.validate.Struct(user)
	if err != nil {
		return utils.HandlerError(err)
	}

	// hash password
	hash := utils.HashPassword(user.Password)
	user.Password = hash

	return s.repoAuth.Register(user)
}

func (s *authService) Login(user *dto.Login) (*dto.LoginResponse, error) {
	// validation
	err := s.validate.Struct(user)
	if err != nil {
		return nil, utils.HandlerError(err)
	}

	// check email
	res, err := s.repoAuth.Login(user.Email)
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
