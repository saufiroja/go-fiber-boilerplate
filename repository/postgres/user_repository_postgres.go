package postgres

import (
	"project/go-fiber-boilerplate/dto"
	"project/go-fiber-boilerplate/entity"
	"project/go-fiber-boilerplate/interfaces"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &repository{
		DB: db,
	}
}

func (repo *repository) FindAllUsers() ([]dto.FindAllUsers, error) {
	data := []dto.FindAllUsers{}

	err := repo.DB.Model(&entity.User{}).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *repository) FindUserByID(id string) (*dto.FindUserByID, error) {
	data := &dto.FindUserByID{}

	err := repo.DB.Model(&entity.User{}).Where("id = ?", id).First(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *repository) UpdateUserByID(id string, user *dto.UpdateUserByID) error {
	data := &entity.User{
		FullName: user.FullName,
		Email:    user.Email,
		IsMale:   user.IsMale,
	}

	err := repo.DB.Model(&entity.User{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) DeleteUserByID(id string) error {
	err := repo.DB.Model(&entity.User{}).Where("id = ?", id).Delete(&entity.User{}).Error
	if err != nil {
		return err
	}

	return nil
}
