package dtos

import "github.com/invopop/validation"

type (
	Profile struct {
		UserName  string `json:"username"`
		Bio       string `json:"bio"`
		Image     string `json:"image"`
		Following bool   `json:"following"`
	}

	ProfileData struct {
		Profile Profile `json:"profile"`
	}

	ProfileRequest struct {
		UserID   int64
		Email    string
		UserName string `param:"username"`
	}
)

func (p ProfileRequest) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.UserName, validation.Required),
	)
}
