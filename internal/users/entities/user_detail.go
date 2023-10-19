package entities

import "github.com/DoWithLogic/go-echo-realworld/internal/users/dtos"

func NewUserDetail(data Users) dtos.UserDetailResponse {
	return dtos.UserDetailResponse{
		Email:    data.Email,
		Token:    "",
		UserName: data.UserName,
		Bio:      data.Bio,
		Image:    data.Image,
	}
}
