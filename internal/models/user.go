package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(connection *gorm.DB) *UserModel {
	return &UserModel{
		db: connection,
	}
}

func (um *UserModel) Login(username string, password string) (User, error) {
	var result User
	err := um.db.Where("username = ? AND password = ?", username, password).First(&result).Error
	if err != nil {
		return User{}, err
	}
	return result, nil
}

func (um *UserModel) Register(newUser User) (bool, error) {
	err := um.db.Create(&newUser).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
