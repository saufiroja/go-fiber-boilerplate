package user

import (
	"project/go-fiber-boilerplate/interfaces"
	"project/go-fiber-boilerplate/models/dto"
	"project/go-fiber-boilerplate/utils/constants"
)

type userService struct {
	repoUser interfaces.UserRepository
	validate constants.IValidation
}

func NewUserService(repoUser interfaces.UserRepository, validate constants.IValidation) interfaces.UserService {
	return &userService{
		repoUser: repoUser,
		validate: validate,
	}
}

func (s *userService) FindAllUsers() ([]dto.FindAllUsers, error) {
	return s.repoUser.FindAllUsers()
}

func (s *userService) FindUserByID(id string) (*dto.FindUserByID, error) {
	res, err := s.repoUser.FindUserByID(id)
	if err != nil {
		return nil, constants.NewBadRequestError("user not found")
	}

	return res, nil
}

func (s *userService) UpdateUserByID(id string, user *dto.UpdateUserByID) error {
	err := s.validate.Validate(user)
	if err != nil {
		return s.validate.ValidationMessage(err)
	}

	data, err := s.FindUserByID(id)
	if err != nil {
		return constants.NewBadRequestError("user not found")
	}

	return s.repoUser.UpdateUserByID(data.ID, user)
}

func (s *userService) DeleteUserByID(id string) error {
	data, err := s.FindUserByID(id)
	if err != nil {
		return constants.NewBadRequestError("user not found")
	}

	return s.repoUser.DeleteUserByID(data.ID)
}
