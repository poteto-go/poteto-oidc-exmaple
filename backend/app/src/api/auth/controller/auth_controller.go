package controller

import (
	"net/http"
	"sync"

	"github.com/poteto-go/poteto"
	"github.com/poteto-go/poteto-hono-oidc/app/src/api/auth/service"
)

type AuthController struct {
	lock    sync.Mutex
	service service.IAuthService
}

type IAuthController interface {
	Login(poteto.Context) error
	TokenRequest(poteto.Context) error
}

func NewAuthController() IAuthController {
	return &AuthController{
		service: service.NewAuthService(),
	}
}

func (ac *AuthController) Login(ctx poteto.Context) error {
	ac.lock.Lock()
	defer ac.lock.Unlock()

	user, err := ac.service.VerifyToken(ctx)
	if err != nil {
		return ctx.JSON(403, map[string]string{
			"auth": "failed",
		})
	}

	return ctx.JSON(200, user)
}

func (ac *AuthController) TokenRequest(ctx poteto.Context) error {
	resp, err := ac.service.TokenRequest(ctx)
	if err == nil {
		return ctx.JSON(http.StatusOK, resp)
	}

	if err.Error() == "BadRequest" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}

	if err.Error() == "BadGateway" {
		return ctx.JSON(http.StatusBadGateway, map[string]string{
			"message": "bad gateway",
		})
	}

	if err.Error() == "Unauthorized" {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{
			"message": "unauthorized",
		})
	}

	return ctx.JSON(http.StatusInternalServerError, map[string]string{
		"message": "server error",
	})
}
