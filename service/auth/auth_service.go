package auth

import (
	"project/go-fiber-boilerplate/interfaces"
	"project/go-fiber-boilerplate/models/dto"
	"project/go-fiber-boilerplate/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()

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
		return nil, utils.NewCustomError(400, "email not found")
	}

	err = s.password.ComparePassword(res.Password, user.Password)
	if err != nil {
		return nil, utils.NewCustomError(400, "password not match")
	}

	accessToken, expiredAccessToken, err := s.token.GenerateAccessToken(res.ID, res.Email, res.FullName)
	if err != nil {
		return nil, utils.NewCustomError(500, "failed to generate access token")
	}

	refreshToken, expiredRefreshToken, err := s.token.GenerateRefreshToken(res.ID, res.Email, res.FullName)
	if err != nil {
		return nil, utils.NewCustomError(500, "failed to generate refresh token")
	}

	token := &dto.LoginResponse{
		AccessToken:         accessToken,
		AccessTokenExpired:  expiredAccessToken,
		RefreshToken:        refreshToken,
		RefreshTokenExpired: expiredRefreshToken,
	}

	return token, nil
}
