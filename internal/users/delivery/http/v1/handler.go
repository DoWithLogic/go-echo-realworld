package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/DoWithLogic/go-echo-realworld/config"
	"github.com/DoWithLogic/go-echo-realworld/internal/users/dtos"
	usecases "github.com/DoWithLogic/go-echo-realworld/internal/users/usecase"
	"github.com/DoWithLogic/go-echo-realworld/pkg/middleware"
	"github.com/DoWithLogic/go-echo-realworld/pkg/otel/zerolog"
	"github.com/DoWithLogic/go-echo-realworld/pkg/utils/response"
	"github.com/labstack/echo/v4"
)

type (
	Handlers interface {
		Login(c echo.Context) error
		CreateUser(c echo.Context) error
		UserDetail(c echo.Context) error
		UpdateUser(c echo.Context) error

		ProfileDetail(c echo.Context) error
		ProfileFollowUser(c echo.Context) error
	}

	handlers struct {
		uc  usecases.Usecase
		log *zerolog.Logger
		cfg config.Config
	}
)

func NewHandlers(uc usecases.Usecase, log *zerolog.Logger) Handlers {
	return &handlers{uc: uc, log: log}
}

func (h *handlers) Login(c echo.Context) error {
	var (
		request dtos.UserData
	)

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(http.StatusBadRequest, response.MsgFailed, err.Error()))
	}

	if err := request.ValidateLogin(); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(http.StatusBadRequest, response.MsgFailed, err.Error()))
	}

	authData, httpCode, err := h.uc.Login(c.Request().Context(), request)
	if err != nil {
		return c.JSON(httpCode, response.NewResponseError(httpCode, response.MsgFailed, err.Error()))
	}

	return c.JSON(httpCode, authData)
}

func (h *handlers) CreateUser(c echo.Context) error {
	var (
		ctx, cancel = context.WithTimeout(c.Request().Context(), time.Duration(30*time.Second))
		request     dtos.UserData
	)
	defer cancel()

	if err := c.Bind(&request); err != nil {
		h.log.Z().Err(err).Msg("users.handlers.CreateUser.Bind")

		return c.JSON(http.StatusBadRequest, response.NewResponseError(
			http.StatusBadRequest,
			response.MsgFailed,
			err.Error()),
		)
	}

	if err := request.ValidateCreate(); err != nil {
		h.log.Z().Err(err).Msg("users.handlers.CreateUser.Validate")

		return c.JSON(http.StatusBadRequest, response.NewResponseError(
			http.StatusBadRequest,
			response.MsgFailed,
			err.Error()),
		)
	}

	data, httpCode, err := h.uc.Create(ctx, request)
	if err != nil {
		return c.JSON(httpCode, response.NewResponseError(
			httpCode,
			response.MsgFailed,
			err.Error()),
		)
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handlers) UserDetail(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(30*time.Second))
	defer cancel()

	userID := c.Get("identity").(*middleware.CustomClaims).UserID

	data, code, err := h.uc.Detail(ctx, userID)
	if err != nil {
		return c.JSON(code, response.NewResponseError(code, response.MsgFailed, err.Error()))
	}

	return c.JSON(code, data)
}

func (h *handlers) UpdateUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(30*time.Second))
	defer cancel()

	var request dtos.UserData
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(http.StatusBadRequest, response.MsgFailed, err.Error()))
	}

	identity := c.Get("identity").(*middleware.CustomClaims)

	data, code, err := h.uc.Update(ctx, request, *identity)
	if err != nil {
		return c.JSON(code, response.NewResponseError(code, response.MsgFailed, err.Error()))
	}

	return c.JSON(code, data)
}

func (h *handlers) ProfileDetail(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(30*time.Second))
	defer cancel()

	var request dtos.ProfileRequest
	if identity := c.Get("identity"); identity != nil {
		request.UserID = identity.(*middleware.CustomClaims).UserID
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(http.StatusBadRequest, response.MsgFailed, err.Error()))
	}

	if err := request.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(http.StatusBadRequest, response.MsgFailed, err.Error()))
	}

	profileData, code, err := h.uc.ProfileDetail(ctx, request)
	if err != nil {
		return c.JSON(code, response.NewResponseError(code, response.MsgFailed, err.Error()))
	}

	return c.JSON(code, profileData)
}

func (h *handlers) ProfileFollowUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(30*time.Second))
	defer cancel()

	var request = dtos.ProfileRequest{
		UserID: c.Get("identity").(*middleware.CustomClaims).UserID,
		Email:  c.Get("identity").(*middleware.CustomClaims).Email,
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(http.StatusBadRequest, response.MsgFailed, err.Error()))
	}

	if err := request.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(http.StatusBadRequest, response.MsgFailed, err.Error()))
	}

	profileData, code, err := h.uc.FollowUser(ctx, request)
	if err != nil {
		return c.JSON(code, response.NewResponseError(code, response.MsgFailed, err.Error()))
	}

	return c.JSON(code, profileData)
}
