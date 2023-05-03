package dto

type Register struct {
	FullName string `json:"full_name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	IsMale   bool   `json:"is_male"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginResponse struct {
	AccessToken         string `json:"access_token"`
	AccessTokenExpired  int64  `json:"access_token_expired"`
	RefreshToken        string `json:"refresh_token"`
	RefreshTokenExpired int64  `json:"refresh_token_expired"`
}

type FindAllUsers struct {
	ID        string `json:"id"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	IsMale    bool   `json:"is_male"`
	CreatedAt int64  `json:"created_at"`
}

type FindUserByID struct {
	ID        string `json:"id"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	IsMale    bool   `json:"is_male"`
	CreatedAt int64  `json:"created_at"`
}

type UpdateUserByID struct {
	FullName string `json:"full_name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	IsMale   bool   `json:"is_male"`
}
