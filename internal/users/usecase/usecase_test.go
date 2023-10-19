package usecase_test

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/DoWithLogic/go-echo-realworld/config"
	"github.com/DoWithLogic/go-echo-realworld/internal/users/dtos"
	"github.com/DoWithLogic/go-echo-realworld/internal/users/entities"
	mocks "github.com/DoWithLogic/go-echo-realworld/internal/users/mock"
	"github.com/DoWithLogic/go-echo-realworld/internal/users/usecase"
	"github.com/DoWithLogic/go-echo-realworld/pkg/apperror"
	"github.com/DoWithLogic/go-echo-realworld/pkg/otel/zerolog"
	"github.com/DoWithLogic/go-echo-realworld/pkg/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_usecase_Detail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	repo := mocks.NewMockRepository(ctrl)
	uc := usecase.NewUseCase(repo, zerolog.NewZeroLog(ctx, os.Stdout), config.Config{})

	var id int64 = 1

	returnedDetail := entities.Users{
		UserID:    id,
		Email:     "test@test.com",
		UserName:  "test",
		Bio:       "",
		Image:     "",
		CreatedAt: time.Now(),
		CreatedBy: "SYSTEM",
	}

	t.Run("detail_positive", func(t *testing.T) {
		repo.EXPECT().GetUserByID(ctx, id).Return(returnedDetail, nil)

		detail, httpCode, err := uc.Detail(ctx, id)
		require.NoError(t, err)
		require.Equal(t, httpCode, http.StatusOK)
		require.Equal(t, detail, entities.NewUserDetail(returnedDetail))
	})

	t.Run("detail_negative_failed_query_detail", func(t *testing.T) {
		repo.EXPECT().GetUserByID(ctx, id).Return(entities.Users{}, sql.ErrNoRows)

		detail, httpCode, err := uc.Detail(ctx, id)
		require.EqualError(t, err, sql.ErrNoRows.Error())
		require.Equal(t, httpCode, http.StatusInternalServerError)
		require.Equal(t, detail, dtos.UserDetailResponse{})
	})

}

func Test_usecase_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		password = "testing"
		email    = "martin@test.com"

		config = config.Config{
			Authentication: config.AuthenticationConfig{
				Key:       "DoWithLogic!@#",
				SecretKey: "s3cr#tK3y!@#v001",
				SaltKey:   "s4ltK3y!@#ddv001",
			},
		}
	)

	ctx := context.Background()
	repo := mocks.NewMockRepository(ctrl)
	uc := usecase.NewUseCase(repo, zerolog.NewZeroLog(ctx, os.Stdout), config)

	returnedUser := entities.Users{
		UserID:   1,
		Email:    email,
		Password: utils.Encrypt(password, config),
	}

	t.Run("login_positive", func(t *testing.T) {
		repo.EXPECT().GetUserByEmail(ctx, email).Return(returnedUser, nil)

		authData, code, err := uc.Login(ctx, dtos.UserLoginRequest{Email: email, Password: password})
		require.NoError(t, err)
		require.Equal(t, code, http.StatusOK)
		require.NotNil(t, authData)

	})

	t.Run("login_negative_invalid_password", func(t *testing.T) {
		repo.EXPECT().GetUserByEmail(ctx, email).Return(returnedUser, nil)

		authData, code, err := uc.Login(ctx, dtos.UserLoginRequest{Email: email, Password: "testingpwd"})
		require.EqualError(t, apperror.ErrInvalidPassword, err.Error())
		require.Equal(t, code, http.StatusUnauthorized)
		require.Equal(t, authData, dtos.UserLoginResponse{})

	})

	t.Run("login_negative_failed_query_email", func(t *testing.T) {
		repo.EXPECT().GetUserByEmail(ctx, email).Return(entities.Users{}, sql.ErrNoRows)

		authData, code, err := uc.Login(ctx, dtos.UserLoginRequest{Email: email, Password: password})
		require.EqualError(t, err, sql.ErrNoRows.Error())
		require.Equal(t, code, http.StatusInternalServerError)
		require.Equal(t, authData, dtos.UserLoginResponse{})

	})

}
