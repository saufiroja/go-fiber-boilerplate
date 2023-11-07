package auth

import (
	"project/go-fiber-boilerplate/interfaces"
	"project/go-fiber-boilerplate/models/dto"
	"project/go-fiber-boilerplate/utils"
	"project/go-fiber-boilerplate/utils/constants"
	"time"

	"github.com/google/uuid"
)

type authService struct {
	repoAuth interfaces.UserRepository
	validate constants.IValidation
	token    *utils.GenerateToken
	password *utils.Password
}

func NewAuthService(
	repoAuth interfaces.UserRepository,
	token *utils.GenerateToken,
	password *utils.Password,
	validate constants.IValidation,
) interfaces.AuthService {
	return &authService{
		repoAuth: repoAuth,
		validate: validate,
		token:    token,
		password: password,
	}
}

func (s *authService) Register(user *dto.Register) error {
	// validation
	err := s.validate.Validate(user)
	if err != nil {
		return s.validate.ValidationMessage(err)
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
	err := s.validate.Validate(user)
	if err != nil {
		return nil, s.validate.ValidationMessage(err)
	}

	// check email
	res, err := s.repoAuth.FindUserByEmail(user.Email)
	if err != nil {
		return nil, constants.NewBadRequestError("email not found")
	}

	err = s.password.ComparePassword(res.Password, user.Password)
	if err != nil {
		return nil, constants.NewBadRequestError("password is wrong")
	}

	accessToken, expiredAccessToken, err := s.token.GenerateAccessToken(res.ID, res.Email, res.FullName)
	if err != nil {
		return nil, constants.NewBadRequestError(err.Error())
	}

	refreshToken, expiredRefreshToken, err := s.token.GenerateRefreshToken(res.ID, res.Email, res.FullName)
	if err != nil {
		return nil, constants.NewBadRequestError(err.Error())
	}

	token := &dto.LoginResponse{
		AccessToken:         accessToken,
		AccessTokenExpired:  expiredAccessToken,
		RefreshToken:        refreshToken,
		RefreshTokenExpired: expiredRefreshToken,
	}

	return token, nil
}
