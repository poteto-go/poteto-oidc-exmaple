package usecase

import (
	"time"

	"github.com/poteto-go/poteto-hono-oidc/app/src/config"
	"github.com/poteto-go/poteto/oidc"
)

// VerifyClaims checks the validity of the OIDC claims.
func VerifyClaims(claims oidc.GoogleOidcClaims) bool {
	// 1. Check ClientId
	if claims.Azp != config.AppConfig.ClientId {
		return false
	}

	// 2. Check Issuer
	if claims.Iss != "https://accounts.google.com" {
		return false
	}

	// 3. Check Expiration
	if claims.Exp < time.Now().Unix() {
		return false
	}

	// 4. Check Email Verified
	if !claims.EmailVerified {
		return false
	}

	return true
}
