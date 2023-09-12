package auth

import (
	"project/go-fiber-boilerplate/interfaces"
	"project/go-fiber-boilerplate/models/dto"
	"project/go-fiber-boilerplate/utils"

	"github.com/go-playground/validator/v10"
)

type authService struct {
	repoAuth interfaces.UserRepository
	validate *validator.Validate
	token    *utils.GenerateToken
	password *utils.Password
}

func NewAuthService(repoAuth interfaces.UserRepository, token *utils.GenerateToken, password *utils.Password) interfaces.AuthService {
	return &authService{
		repoAuth: repoAuth,
		validate: validator.New(),
		token:    token,
		password: password,
	}
}

func (s *authService) Register(user *dto.Register) error {
	// validation
	err := s.validate.Struct(user)
	if err != nil {
		return utils.HandlerError(err)
	}

	// hash password
	hash := s.password.HashPassword(user.Password)
	user.Password = hash

	return s.repoAuth.InsertUser(user)
}

func (s *authService) Login(user *dto.Login) (*dto.LoginResponse, error) {
	// validation
	err := s.validate.Struct(user)
	if err != nil {
		return nil, utils.HandlerError(err)
	}

	// check email
	res, err := s.repoAuth.FindUserByEmail(user.Email)
	if err != nil {
		return nil, utils.HandlerErrorCustom(404, "email not found")
	}

	err = s.password.ComparePassword(res.Password, user.Password)
	if err != nil {
		return nil, utils.HandlerErrorCustom(400, "wrong password")
	}

	accessToken, expiredAccessToken, err := s.token.GenerateAccessToken(res.ID, res.Email, res.FullName)
	if err != nil {
		return nil, utils.HandlerErrorCustom(500, "failed to generate access token")
	}

	refreshToken, expiredRefreshToken, err := s.token.GenerateRefreshToken(res.ID, res.Email, res.FullName)
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
