package dtos

import (
	"github.com/invopop/validation"
)

type CreateUserRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	UserName string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

func (cup CreateUserRequest) Validate() error {
	return validation.ValidateStruct(&cup,
		validation.Field(&cup.UserName, validation.Required),
		validation.Field(&cup.Email, validation.Required),
		validation.Field(&cup.Password, validation.Required),
	)
}
