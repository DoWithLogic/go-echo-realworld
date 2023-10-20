package dtos

import "github.com/invopop/validation"

type (
	UserRequest struct {
		Data User `json:"user"`
	}
	UserResponse struct {
		Data User `json:"user"`
	}

	User struct {
		UserName string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password,omitempty"`
		Token    string `json:"token"`
		Image    string `json:"image"`
		Bio      string `json:"bio,"`
	}
)

func (x UserRequest) ValidateCreate() error {
	return validation.ValidateStruct(&x.Data,
		validation.Field(&x.Data.UserName, validation.Required),
		validation.Field(&x.Data.Email, validation.Required),
		validation.Field(&x.Data.Password, validation.Required),
	)
}

func (x UserRequest) ValidateLogin() error {
	return validation.ValidateStruct(&x.Data,
		validation.Field(&x.Data.Email, validation.Required),
		validation.Field(&x.Data.Password, validation.Required),
	)
}
