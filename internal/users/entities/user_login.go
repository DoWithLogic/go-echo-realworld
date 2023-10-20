package entities

import (
	"github.com/DoWithLogic/go-echo-realworld/internal/users/dtos"
)

func NewUserLogin(res Users, token string) dtos.UserResponse {
	return dtos.UserResponse{
		Data: dtos.User{
			Email: res.Email,
			Token: token,
			Bio:   res.Bio,
			Image: res.Image,
		},
	}
}
