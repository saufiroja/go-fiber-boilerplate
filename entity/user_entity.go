package entity

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsMale    bool   `json:"is_male"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at" gorm:"default:null"`
}
