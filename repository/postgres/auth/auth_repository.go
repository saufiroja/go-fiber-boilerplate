package auth

import (
	"project/go-fiber-boilerplate/interfaces"
	"project/go-fiber-boilerplate/models/dto"
	"project/go-fiber-boilerplate/models/entity"

	"gorm.io/gorm"
)

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) interfaces.AuthRepository {
	return &authRepository{
		DB: db,
	}
}

func (r *authRepository) Register(user *dto.Register) error {
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

func (r *authRepository) Login(email string) (*entity.User, error) {
	data := &entity.User{}

	err := r.DB.Model(&entity.User{}).Where("email = ?", email).First(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
