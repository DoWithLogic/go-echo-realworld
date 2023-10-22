package dtos

import (
	"github.com/DoWithLogic/go-echo-realworld/pkg/apperror"
	"github.com/DoWithLogic/go-echo-realworld/pkg/constant"
	"github.com/invopop/validation"
)

type (
	UserData struct {
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

	UpdateUserStatusRequest struct {
		UserID   int64  `json:"-"`
		UpdateBy string `json:"-"`
		Status   int    `json:"status"`
	}

	UpdateUserRequest struct {
		UserID      int64  `json:"-"`
		Fullname    string `json:"fullname"`
		PhoneNumber string `json:"phone_number"`
		UserType    string `json:"user_type"`
		Email       string `json:"email"`
		Password    string `json:"password"`
		UpdateBy    string `json:"-"`
	}
)

func (x UserData) ValidateCreate() error {
	return validation.ValidateStruct(&x.Data,
		validation.Field(&x.Data.UserName, validation.Required),
		validation.Field(&x.Data.Email, validation.Required),
		validation.Field(&x.Data.Password, validation.Required),
	)
}

func (x UserData) ValidateLogin() error {
	return validation.ValidateStruct(&x.Data,
		validation.Field(&x.Data.Email, validation.Required),
		validation.Field(&x.Data.Password, validation.Required),
	)
}

func (ussp UpdateUserStatusRequest) Validate() error {
	return validation.ValidateStruct(&ussp,
		validation.Field(&ussp.UserID, validation.NotNil),
		validation.Field(&ussp.UpdateBy, validation.NotNil),
		validation.Field(&ussp.Status, validation.Required),
	)
}

func (cup UpdateUserRequest) Validate() error {
	if cup.UserType != "" && cup.UserType != constant.UserTypePremium && cup.UserType != constant.UserTypeRegular {
		return apperror.ErrInvalidUserType
	}

	return validation.ValidateStruct(&cup,
		validation.Field(&cup.UserID, validation.NotNil),
		validation.Field(&cup.UpdateBy, validation.NotNil),
	)
}
