package entities

import "github.com/DoWithLogic/go-echo-realworld/internal/users/dtos"

func NewUserDetail(data Users) dtos.UserResponse {
	return dtos.UserResponse{
		Data: dtos.User{
			Email:    data.Email,
			UserName: data.UserName,
			Bio:      data.Bio,
			Image:    data.Image,
		},
	}
}
