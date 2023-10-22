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
	"github.com/DoWithLogic/go-echo-realworld/pkg/datasource"
	"github.com/DoWithLogic/go-echo-realworld/pkg/middleware"
	"github.com/DoWithLogic/go-echo-realworld/pkg/otel/zerolog"
	"github.com/DoWithLogic/go-echo-realworld/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type (
	Usecase interface {
		Login(ctx context.Context, request dtos.UserData) (response dtos.UserData, httpCode int, err error)
		Create(ctx context.Context, request dtos.UserData) (response dtos.UserData, httpCode int, err error)
		Detail(ctx context.Context, id int64) (response dtos.UserData, httpCode int, err error)
		Update(ctx context.Context, request dtos.UserData, identity middleware.CustomClaims) (response dtos.UserData, httpCode int, err error)

		FollowUser(ctx context.Context, req dtos.ProfileRequest) (dtos.ProfileData, int, error)
		ProfileDetail(ctx context.Context, req dtos.ProfileRequest) (dtos.ProfileData, int, error)
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

func (uc *usecase) Login(ctx context.Context, request dtos.UserData) (response dtos.UserData, httpCode int, err error) {
	dataLogin, err := uc.repo.GetUserByEmail(ctx, request.Data.Email)
	if err != nil {
		if errors.Is(err, datasource.ErrDataNotFound) {
			return response, http.StatusBadRequest, err
		}

		return response, http.StatusInternalServerError, err
	}

	decryptedPassword, err := utils.Decrypt(dataLogin.Password, uc.cfg)
	if err != nil {
		return response, http.StatusInternalServerError, err
	}

	if !strings.EqualFold(decryptedPassword, request.Data.Password) {
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

func (uc *usecase) Create(ctx context.Context, req dtos.UserData) (res dtos.UserData, httpCode int, err error) {
	if exist := uc.repo.IsUserExist(ctx, req.Data.Email); exist {
		return res, http.StatusConflict, apperror.ErrEmailAlreadyExist
	}

	req.Data.Password, err = utils.Encrypt(req.Data.Password, uc.cfg)
	if err != nil {
		return res, http.StatusBadRequest, err
	}

	if _, err = uc.repo.SaveNewUser(ctx, entities.NewCreateUser(req)); err != nil {
		uc.log.Z().Err(err).Msg("users.uc.Create.SaveNewUser")

		return res, http.StatusInternalServerError, err
	}

	res.Data = dtos.User{
		Email:    req.Data.Email,
		UserName: req.Data.UserName,
	}

	return res, http.StatusOK, nil
}

func (uc *usecase) Detail(ctx context.Context, id int64) (detail dtos.UserData, httpCode int, err error) {
	userDetail, err := uc.repo.GetUserByID(ctx, id)
	if err != nil {
		return detail, http.StatusInternalServerError, err
	}

	return entities.NewUserDetail(userDetail), http.StatusOK, nil
}

func (uc *usecase) Update(
	/*req*/ ctx context.Context, request dtos.UserData, identity middleware.CustomClaims) (
	/*res*/ response dtos.UserData, httpCode int, err error,
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

func (uc *usecase) FollowUser(ctx context.Context, req dtos.ProfileRequest) (dtos.ProfileData, int, error) {
	profile, err := uc.repo.GetUserProfile(ctx, req.UserName)
	if err != nil {
		return dtos.ProfileData{}, http.StatusInternalServerError, err
	}

	profileData := entities.NewProfileDetail(profile)
	_, err = uc.repo.SaveNewProfile(ctx, entities.NewStoreProfile(profile, req))
	if err != nil {
		return dtos.ProfileData{}, http.StatusInternalServerError, err
	}

	profileData.Profile.Following = true

	return profileData, http.StatusOK, nil
}

func (uc *usecase) ProfileDetail(ctx context.Context, req dtos.ProfileRequest) (dtos.ProfileData, int, error) {
	profile, err := uc.repo.GetUserProfile(ctx, req.UserName)
	if err != nil {
		return dtos.ProfileData{}, http.StatusInternalServerError, err
	}

	profileData := entities.NewProfileDetail(profile)
	if req.UserID != 0 {
		if follow := uc.repo.IsUserFollowed(ctx, req.UserID, profile.UserID); follow {
			profileData.Profile.Following = true
		}
	}

	return profileData, http.StatusOK, err
}
