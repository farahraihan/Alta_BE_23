package users

import "restAPi/internal/models"

type LoginResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Token    string `json:"token"`
}

func ToLoginResponse(input models.User, token string) LoginResponse {
	return LoginResponse{
		ID:       input.ID,
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
		Token:    token,
	}
}
