package v1

import (
	"github.com/DoWithLogic/go-echo-realworld/config"
	"github.com/DoWithLogic/go-echo-realworld/pkg/middleware"
	"github.com/labstack/echo/v4"
)

func MapUserRoute(version *echo.Group, h Handlers, cfg config.Config) {
	users := version.Group("users")
	users.POST("", h.CreateUser)
	users.POST("/login", h.Login)
	users.GET("/detail", h.UserDetail, middleware.AuthorizeJWT(cfg))
}
