package user

import (
	"errors"
	"project/go-fiber-boilerplate/dto"
	"project/go-fiber-boilerplate/interfaces"
	"project/go-fiber-boilerplate/utils"

	"github.com/go-playground/validator/v10"
)

type service struct {
	repo     interfaces.UserRepository
	validate *validator.Validate
}

func NewUserService(repo interfaces.UserRepository) interfaces.UserService {
	return &service{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *service) FindAllUsers() ([]dto.FindAllUsers, error) {
	return s.repo.FindAllUsers()
}

func (s *service) FindUserByID(id string) (*dto.FindUserByID, error) {
	return s.repo.FindUserByID(id)
}

func (s *service) UpdateUserByID(id string, user *dto.UpdateUserByID) error {
	err := s.validate.Struct(user)
	if err != nil {
		return utils.HandlerError(err)
	}

	data, err := s.FindUserByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	return s.repo.UpdateUserByID(data.ID, user)
}

func (s *service) DeleteUserByID(id string) error {
	data, err := s.FindUserByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	return s.repo.DeleteUserByID(data.ID)
}
