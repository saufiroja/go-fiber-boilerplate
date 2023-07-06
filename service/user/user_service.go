package user

import (
	"errors"
	"project/go-fiber-boilerplate/interfaces"
	"project/go-fiber-boilerplate/models/dto"
	"project/go-fiber-boilerplate/utils"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	repoUser interfaces.UserRepository
	validate *validator.Validate
}

func NewUserService(repoUser interfaces.UserRepository) interfaces.UserService {
	return &userService{
		repoUser: repoUser,
		validate: validator.New(),
	}
}

func (s *userService) FindAllUsers() ([]dto.FindAllUsers, error) {
	return s.repoUser.FindAllUsers()
}

func (s *userService) FindUserByID(id string) (*dto.FindUserByID, error) {
	return s.repoUser.FindUserByID(id)
}

func (s *userService) UpdateUserByID(id string, user *dto.UpdateUserByID) error {
	err := s.validate.Struct(user)
	if err != nil {
		return utils.HandlerError(err)
	}

	data, err := s.FindUserByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	return s.repoUser.UpdateUserByID(data.ID, user)
}

func (s *userService) DeleteUserByID(id string) error {
	data, err := s.FindUserByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	return s.repoUser.DeleteUserByID(data.ID)
}
