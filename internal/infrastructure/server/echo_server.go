package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ungpa/goon-ah-lang/internal/infrastructure/config"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewEchoServer),
)

func NewEchoServer() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderXRequestedWith, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "alive")
	})

	return e
}

func StartServer(lifecycle fx.Lifecycle, config *config.AppConfig, e *echo.Echo) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				e.Logger.Fatal(e.Start(":" + config.Port))
			}()
			return nil
		},
	})
}
