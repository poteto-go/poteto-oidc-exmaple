package app

import (
	"net/http"

	"github.com/poteto-go/poteto"
	"github.com/poteto-go/poteto-hono-oidc/app/src/api/auth"
	"github.com/poteto-go/poteto-hono-oidc/app/src/config"
	"github.com/poteto-go/poteto/middleware"
)

func NewApp() poteto.Poteto {
	p := poteto.New()

	// CORS
	p.Register(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Security Header
	p.Register(middleware.CamaraWithConfig(middleware.DefaultCamaraConfig))

	// Request Logger
	p.Register(
		middleware.RequestLoggerWithConfig(
			config.NewRequestLoggerConfig(),
		),
	)

	// Timeout
	p.Register(
		middleware.TimeoutWithConfig(
			middleware.DefaultTimeoutConfig, // 10s
		),
	)

	p.GET("/", func(ctx poteto.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{
			"hello": "world",
		})
	})

	authApiV1 := auth.NewAuthApi()
	p.AddApi(authApiV1)

	return p
}
