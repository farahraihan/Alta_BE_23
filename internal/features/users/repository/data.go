package repository

import (
	"clean/internal/features/todos/repository"
	"clean/internal/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"hp"`
	// Alamats  string
	Todos []repository.Todo `gorm:"foreignKey:owner"`
}

func (u *User) toUserEntity() users.User {
	return users.User{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Phone:    u.Phone,
	}
}

func toUserData(input users.User) User {
	return User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	}
}
