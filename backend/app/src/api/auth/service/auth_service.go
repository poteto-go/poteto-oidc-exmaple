package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/poteto-go/poteto"
	"github.com/poteto-go/poteto-hono-oidc/app/src/api/auth/schema"
	"github.com/poteto-go/poteto-hono-oidc/app/src/api/auth/service/usecase"
	"github.com/poteto-go/poteto-hono-oidc/app/src/config"
	"github.com/poteto-go/poteto-hono-oidc/app/src/domains/domain"
	"github.com/poteto-go/poteto/oidc"
)

type AuthService struct{}

type IAuthService interface {
	VerifyToken(poteto.Context) (domain.User, error)
	TokenRequest(poteto.Context) (domain.TokenResponse, error)
}

func NewAuthService() IAuthService {
	return &AuthService{}
}

func (s *AuthService) VerifyToken(ctx poteto.Context) (domain.User, error) {
	token, ok := ctx.Get("googleToken")
	if !ok {
		return domain.User{}, errors.New("google token not found in context")
	}

	claims := oidc.GoogleOidcClaims{}
	err := json.Unmarshal(token.([]byte), &claims)
	if err != nil {
		return domain.User{}, errors.New("failed to unmarshal google token claims")
	}

	// Verify claims using the usecase
	if !usecase.VerifyClaims(claims) {
		return domain.User{}, errors.New("invalid verify: claims verification failed")
	}

	user := domain.User{
		Email:   claims.Email,
		Name:    claims.Name,
		Picture: claims.Picture,
	}

	return user, nil
}

func (s *AuthService) TokenRequest(ctx poteto.Context) (domain.TokenResponse, error) {
	code, ok := ctx.QueryParam("code")
	if !ok {
		return domain.TokenResponse{}, errors.New("BadRequest")
	}

	reqBody := schema.TokenRequestBody{
		Code:         code,
		ClientId:     config.AppConfig.ClientId,
		ClientSecret: config.AppConfig.ClientSecret,
		RedirectUri:  config.AppConfig.RedirectURL,
		GrantType:    "authorization_code",
	}
	reqBodyPayload, err := json.Marshal(reqBody)
	if err != nil {
		return domain.TokenResponse{}, errors.New("BadRequest")
	}

	req, _ := http.NewRequest(
		http.MethodPost,
		config.AppConfig.TokenEndpoint,
		bytes.NewBuffer(reqBodyPayload),
	)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return domain.TokenResponse{}, errors.New("BadGateway")
	}
	defer resp.Body.Close()

	tokenResponse := domain.TokenResponse{}
	rawToken, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.TokenResponse{}, errors.New("Unauthorized")
	}
	json.Unmarshal(rawToken, &tokenResponse)
	return tokenResponse, nil
}
