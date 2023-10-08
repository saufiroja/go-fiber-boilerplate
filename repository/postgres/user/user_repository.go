package user

import (
	"context"
	"database/sql"
	"project/go-fiber-boilerplate/interfaces"
	"project/go-fiber-boilerplate/models/dto"
	"project/go-fiber-boilerplate/models/entity"
	"time"
)

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) interfaces.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) FindAllUsers() ([]dto.FindAllUsers, error) {
	data := []dto.FindAllUsers{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := "SELECT id, full_name, email, is_male, created_at FROM users"

	row, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var user dto.FindAllUsers
		err := row.Scan(&user.ID, &user.FullName, &user.Email, &user.IsMale, &user.CreatedAt)
		if err != nil {
			return nil, err
		}

		data = append(data, user)
	}

	return data, nil
}

func (r *userRepository) FindUserByID(id string) (*dto.FindUserByID, error) {
	data := &dto.FindUserByID{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "SELECT id, full_name, email, is_male FROM users WHERE id = $1"

	row := r.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&data.ID, &data.FullName, &data.Email, &data.IsMale)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *userRepository) UpdateUserByID(id string, user *dto.UpdateUserByID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	trx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	query := "UPDATE users SET full_name = $1, email = $2, is_male = $3, updated_at = $4 WHERE id = $5"
	_, err = trx.ExecContext(ctx, query, user.FullName, user.Email, user.IsMale, user.UpdatedAt, id)

	if err != nil {
		return err
	}

	return trx.Commit()
}

func (r *userRepository) DeleteUserByID(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	trx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	query := "DELETE FROM users WHERE id = $1"
	_, err = trx.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return trx.Commit()
}

func (r *userRepository) InsertUser(user *dto.Register) error {
	data := &entity.User{
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
		Password: user.Password,
		IsMale:   user.IsMale,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	trx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	query := "INSERT INTO users (id, email, full_name, password, is_male, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err = trx.ExecContext(ctx, query, data.ID, data.Email, data.FullName, data.Password, data.IsMale, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		return err
	}

	return trx.Commit()
}

func (r *userRepository) FindUserByEmail(email string) (*entity.User, error) {
	data := &entity.User{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "SELECT id, email, full_name, password FROM users WHERE email = $1"

	err := r.DB.QueryRowContext(ctx, query, email).Scan(&data.ID, &data.Email, &data.FullName, &data.Password)
	if err != nil {
		return nil, err
	}

	return data, nil
}
