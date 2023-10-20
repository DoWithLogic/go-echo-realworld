package entities

import (
	"time"

	"github.com/DoWithLogic/go-echo-realworld/config"
	"github.com/DoWithLogic/go-echo-realworld/internal/users/dtos"
	"github.com/DoWithLogic/go-echo-realworld/pkg/middleware"
	"github.com/DoWithLogic/go-echo-realworld/pkg/utils"
)

type CreateUser struct {
	FullName    string
	PhoneNumber string
	UserType    string
	IsActive    bool
	CreatedAt   time.Time
	CreatedBy   string
}

func NewCreateUser(data dtos.UserRequest, cfg config.Config) Users {
	return Users{
		UserName:  data.Data.UserName,
		Email:     data.Data.Email,
		Password:  utils.Encrypt(data.Data.Password, cfg),
		CreatedAt: time.Now(),
		CreatedBy: "SYSTEM",
	}
}

func NewUpdateUser(req dtos.UserRequest, cfg config.Config, identity middleware.CustomClaims) Users {
	return Users{
		UserID:    identity.UserID,
		UserName:  req.Data.UserName,
		Email:     req.Data.Email,
		Password:  req.Data.Password,
		Bio:       req.Data.Bio,
		Image:     req.Data.Image,
		UpdatedAt: time.Now(),
		UpdatedBy: identity.Email,
	}
}
