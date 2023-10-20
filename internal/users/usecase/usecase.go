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
		Login(ctx context.Context, request dtos.UserRequest) (response dtos.UserResponse, httpCode int, err error)
		Create(ctx context.Context, request dtos.UserRequest) (response dtos.UserResponse, httpCode int, err error)
		Detail(ctx context.Context, id int64) (response dtos.UserResponse, httpCode int, err error)
		Update(ctx context.Context, request dtos.UserRequest, identity middleware.CustomClaims) (response dtos.UserResponse, httpCode int, err error)
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

func (uc *usecase) Login(ctx context.Context, request dtos.UserRequest) (response dtos.UserResponse, httpCode int, err error) {
	dataLogin, err := uc.repo.GetUserByEmail(ctx, request.Data.Email)
	if err != nil {
		return response, http.StatusInternalServerError, err
	}

	if !strings.EqualFold(utils.Decrypt(dataLogin.Password, uc.cfg), request.Data.Password) {
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

func (uc *usecase) Create(ctx context.Context, req dtos.UserRequest) (res dtos.UserResponse, httpCode int, err error) {
	if exist := uc.repo.IsUserExist(ctx, req.Data.Email); exist {
		return res, http.StatusConflict, apperror.ErrEmailAlreadyExist
	}

	if _, err = uc.repo.SaveNewUser(ctx, entities.NewCreateUser(req, uc.cfg)); err != nil {
		uc.log.Z().Err(err).Msg("users.uc.Create.SaveNewUser")

		return res, http.StatusInternalServerError, err
	}

	res.Data = dtos.User{
		Email:    req.Data.Email,
		UserName: req.Data.UserName,
	}

	return res, http.StatusOK, nil
}

func (uc *usecase) Detail(ctx context.Context, id int64) (detail dtos.UserResponse, httpCode int, err error) {
	userDetail, err := uc.repo.GetUserByID(ctx, id)
	if err != nil {
		return detail, http.StatusInternalServerError, err
	}

	return entities.NewUserDetail(userDetail), http.StatusOK, nil
}

func (uc *usecase) Update(
	/*req*/ ctx context.Context, request dtos.UserRequest, identity middleware.CustomClaims) (
	/*res*/ response dtos.UserResponse, httpCode int, err error,
) {
	if err = uc.repo.UpdateUser(ctx, entities.NewUpdateUser(request, uc.cfg, identity)); err != nil {
		return response, http.StatusInternalServerError, err
	}

	detail, err := uc.repo.GetUserByID(ctx, identity.UserID)
	if err != nil {
		return response, http.StatusInternalServerError, err
	}

	response.Data = dtos.User{
		UserName: detail.UserName,
		Email:    detail.Email,
		Image:    detail.Image,
		Bio:      detail.Bio,
	}

	return response, http.StatusOK, nil
}
