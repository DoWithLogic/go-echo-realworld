package entities

import (
	"github.com/DoWithLogic/go-echo-realworld/internal/users/dtos"
)

func NewUserLogin(data Users, token string) dtos.UserLoginResponse {
	return dtos.UserLoginResponse{
		Email:    data.Email,
		Token:    "",
		UserName: data.UserName,
		Bio:      data.Bio,
		Image:    data.Image,
	}
}
