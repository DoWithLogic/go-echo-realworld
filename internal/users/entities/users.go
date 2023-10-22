package entities

import (
	"time"

	"github.com/DoWithLogic/go-echo-realworld/config"
	"github.com/DoWithLogic/go-echo-realworld/internal/users/dtos"
	"github.com/DoWithLogic/go-echo-realworld/pkg/middleware"
	"github.com/DoWithLogic/go-echo-realworld/pkg/utils"
)

type (
	Users struct {
		UserID    int64
		Email     string
		Password  string
		UserName  string
		Bio       string
		Image     string
		CreatedAt time.Time
		CreatedBy string
		UpdatedAt time.Time
		UpdatedBy string
	}

	LockingOpt struct {
		PessimisticLocking bool
	}
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
