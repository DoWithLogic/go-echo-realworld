package v1

import (
	"github.com/DoWithLogic/go-echo-realworld/config"
	"github.com/DoWithLogic/go-echo-realworld/pkg/middleware"
	"github.com/labstack/echo/v4"
)

func MapUserRoute(version *echo.Group, h Handlers, cfg config.Config) {
	version.POST("users", h.CreateUser)
	version.POST("users/login", h.Login)
	version.GET("user", h.UserDetail, middleware.AuthorizeJWT(cfg))
	version.PUT("user", h.UpdateUser, middleware.AuthorizeJWT(cfg))

	version.GET("profiles/:username", h.ProfileDetail, middleware.OptionalAuthJWT(cfg))
	version.POST("profiles/:username/follow", h.ProfileFollowUser, middleware.AuthorizeJWT(cfg))
}
