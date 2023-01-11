package auth

import (
	"project/go-fiber-boilerplate/dto"
	"project/go-fiber-boilerplate/entity"
	"project/go-fiber-boilerplate/interfaces"

	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) interfaces.UserRepository {
	return &AuthRepository{
		DB: db,
	}
}

func (repo *AuthRepository) Register(user *dto.Register) error {
	data := &entity.User{
		Email:    user.Email,
		FullName: user.FullName,
		Password: user.Password,
		IsMale:   user.IsMale,
	}

	err := repo.DB.Model(&entity.User{}).Create(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *AuthRepository) Login(email string) (*entity.User, error) {
	data := &entity.User{}

	err := repo.DB.Model(&entity.User{}).Where("email = ?", email).First(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
