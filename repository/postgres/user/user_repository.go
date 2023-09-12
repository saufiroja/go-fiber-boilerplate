package user

import (
	"project/go-fiber-boilerplate/interfaces"
	"project/go-fiber-boilerplate/models/dto"
	"project/go-fiber-boilerplate/models/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) FindAllUsers() ([]dto.FindAllUsers, error) {
	data := []dto.FindAllUsers{}

	err := r.DB.Model(&entity.User{}).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *userRepository) FindUserByID(id string) (*dto.FindUserByID, error) {
	data := &dto.FindUserByID{}

	err := r.DB.Model(&entity.User{}).Where("id = ?", id).First(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *userRepository) UpdateUserByID(id string, user *dto.UpdateUserByID) error {
	data := &entity.User{
		FullName: user.FullName,
		Email:    user.Email,
		IsMale:   user.IsMale,
	}

	err := r.DB.Model(&entity.User{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUserByID(id string) error {
	err := r.DB.Model(&entity.User{}).Where("id = ?", id).Delete(&entity.User{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) InsertUser(user *dto.Register) error {
	data := &entity.User{
		Email:    user.Email,
		FullName: user.FullName,
		Password: user.Password,
		IsMale:   user.IsMale,
	}

	err := r.DB.Model(&entity.User{}).Create(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindUserByEmail(email string) (*entity.User, error) {
	data := &entity.User{}

	err := r.DB.Model(&entity.User{}).Where("email = ?", email).First(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
