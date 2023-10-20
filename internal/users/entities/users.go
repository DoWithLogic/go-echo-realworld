package entities

import (
	"time"
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
