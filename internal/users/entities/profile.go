package entities

import (
	"time"

	"github.com/DoWithLogic/go-echo-realworld/internal/users/dtos"
)

type (
	Profile struct {
		ID           int64
		UserID       int64
		FollowUserID int64
		UserName     string
		Bio          string
		Image        string
		IsActive     bool
		CreatedAt    time.Time
		CreatedBy    string
		UpdatedAt    time.Time
		UpdatedBy    string
	}
)

func NewProfileDetail(p Profile) dtos.ProfileData {
	return dtos.ProfileData{
		Profile: dtos.Profile{
			UserName: p.UserName,
			Bio:      p.Bio,
			Image:    p.Image,
		},
	}
}

func NewStoreProfile(p Profile, pd dtos.ProfileRequest) Profile {
	return Profile{
		UserID:       pd.UserID,
		FollowUserID: p.UserID,
		IsActive:     true,
		CreatedAt:    time.Now(),
		CreatedBy:    pd.Email,
	}
}
