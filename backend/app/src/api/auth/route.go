package auth

import (
	"github.com/poteto-go/poteto"
	"github.com/poteto-go/poteto-hono-oidc/app/src/api/auth/controller"
	"github.com/poteto-go/poteto-hono-oidc/app/src/config"
	"github.com/poteto-go/poteto/middleware"
	"github.com/poteto-go/poteto/oidc"
)

func NewAuthApi() poteto.Poteto {
	oidcConfig := middleware.OidcConfig{
		Idp:                        "google",
		ContextKey:                 "googleToken",
		JwksUrl:                    config.AppConfig.JWKsUrl,
		CustomVerifyTokenSignature: oidc.DefaultVerifyTokenSignature,
	}

	authController := controller.NewAuthController()
	authApi := poteto.Api("/v1/auth", func(api poteto.Leaf) {
		api.Register(
			middleware.OidcWithConfig(
				oidcConfig,
			),
		)

		api.POST("/login", authController.Login)
	})

	authApi.GET("/v1/token_request", authController.TokenRequest)

	return authApi
}
