package entities

import (
	"time"

	"github.com/DoWithLogic/go-echo-realworld/config"
	"github.com/DoWithLogic/go-echo-realworld/internal/users/dtos"
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

func NewCreateUser(data dtos.CreateUserRequest, cfg config.Config) Users {
	return Users{
		UserName:  data.UserName,
		Email:     data.Email,
		Password:  utils.Encrypt(data.Password, cfg),
		CreatedAt: time.Now(),
		CreatedBy: "SYSTEM",
	}
}
