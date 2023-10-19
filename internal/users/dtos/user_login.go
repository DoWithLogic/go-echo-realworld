package dtos

import (
	"github.com/invopop/validation"
)

type (
	UserLoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UserLoginResponse struct {
		Email    string `json:"email"`
		Token    string `json:"token"`
		UserName string `json:"username"`
		Bio      string `json:"bio"`
		Image    string `json:"image"`
	}
)

func (ulr UserLoginRequest) Validate() error {
	return validation.ValidateStruct(&ulr,
		validation.Field(&ulr.Email, validation.Required),
		validation.Field(&ulr.Password, validation.Required),
	)
}
