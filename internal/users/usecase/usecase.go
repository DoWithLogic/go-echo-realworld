package usecase

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/DoWithLogic/go-echo-realworld/config"
	"github.com/DoWithLogic/go-echo-realworld/internal/users/dtos"
	"github.com/DoWithLogic/go-echo-realworld/internal/users/entities"
	"github.com/DoWithLogic/go-echo-realworld/internal/users/repository"
	"github.com/DoWithLogic/go-echo-realworld/pkg/apperror"
	"github.com/DoWithLogic/go-echo-realworld/pkg/middleware"
	"github.com/DoWithLogic/go-echo-realworld/pkg/otel/zerolog"
	"github.com/DoWithLogic/go-echo-realworld/pkg/utils"
	"github.com/dgrijalva/jwt-go"
)

type (
	Usecase interface {
		Login(ctx context.Context, request dtos.UserLoginRequest) (response dtos.UserLoginResponse, httpCode int, err error)
		Create(ctx context.Context, payload dtos.CreateUserRequest) (userID int64, httpCode int, err error)
		Detail(ctx context.Context, id int64) (detail dtos.UserDetailResponse, httpCode int, err error)
	}

	usecase struct {
		repo repository.Repository
		log  *zerolog.Logger
		cfg  config.Config
	}
)

func NewUseCase(repo repository.Repository, log *zerolog.Logger, cfg config.Config) Usecase {
	return &usecase{repo, log, cfg}
}

func (uc *usecase) Login(ctx context.Context, request dtos.UserLoginRequest) (response dtos.UserLoginResponse, httpCode int, err error) {
	dataLogin, err := uc.repo.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return response, http.StatusInternalServerError, err
	}

	if !strings.EqualFold(utils.Decrypt(dataLogin.Password, uc.cfg), request.Password) {
		return response, http.StatusUnauthorized, apperror.ErrInvalidPassword
	}

	identityData := middleware.CustomClaims{
		UserID: dataLogin.UserID,
		Email:  dataLogin.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	token, err := middleware.GenerateJWT(identityData, uc.cfg.Authentication.Key)
	if err != nil {
		return response, http.StatusInternalServerError, apperror.ErrFailedGenerateJWT
	}

	return entities.NewUserLogin(dataLogin, token), http.StatusOK, nil
}

func (uc *usecase) Create(ctx context.Context, payload dtos.CreateUserRequest) (userID int64, httpCode int, err error) {
	if exist := uc.repo.IsUserExist(ctx, payload.Email); exist {
		return userID, http.StatusConflict, apperror.ErrEmailAlreadyExist
	}

	userID, err = uc.repo.SaveNewUser(ctx, entities.NewCreateUser(payload, uc.cfg))
	if err != nil {
		uc.log.Z().Err(err).Msg("users.uc.Create.SaveNewUser")

		return userID, http.StatusInternalServerError, err
	}

	return userID, http.StatusOK, nil
}

func (uc *usecase) Detail(ctx context.Context, id int64) (detail dtos.UserDetailResponse, httpCode int, err error) {
	userDetail, err := uc.repo.GetUserByID(ctx, id)
	if err != nil {
		return detail, http.StatusInternalServerError, err
	}

	return entities.NewUserDetail(userDetail), http.StatusOK, nil
}
