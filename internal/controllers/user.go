package controllers

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"taskday5.2/internal/models"
)

type UserController struct {
	model *models.UserModel
}

func NewUserController(m *models.UserModel) *UserController {
	return &UserController{
		model: m,
	}
}

func (uc *UserController) Login() (models.User, error) {
	reader := bufio.NewReader(os.Stdin)
	var email, password string

	fmt.Print("Masukkan email: ")
	email, _ = reader.ReadString('\n')
	fmt.Print("Masukkan password: ")
	password, _ = reader.ReadString('\n')

	result, err := uc.model.Login(email[:len(email)-1], password[:len(password)-1])
	if err != nil {
		return models.User{}, errors.New("terjadi masalah ketika login")
	}
	return result, nil
}

func (uc *UserController) Register() (models.User, error) {
	reader := bufio.NewReader(os.Stdin)
	var newData models.User

	fmt.Print("Masukkan Nama: ")
	newData.Name, _ = reader.ReadString('\n')
	fmt.Print("Masukkan Password: ")
	newData.Password, _ = reader.ReadString('\n')
	fmt.Print("Masukkan Email: ")
	newData.Email, _ = reader.ReadString('\n')
	fmt.Print("Masukkan HP: ")
	newData.Phone, _ = reader.ReadString('\n')

	result, err := uc.model.Register(models.User{
		Name:     newData.Name[:len(newData.Name)-1],
		Password: newData.Password[:len(newData.Password)-1],
		Email:    newData.Email[:len(newData.Email)-1],
		Phone:    newData.Phone[:len(newData.Phone)-1],
	})
	if err != nil && !result {
		return models.User{}, errors.New("terjadi masalah ketika registrasi")
	}
	return newData, nil
}
