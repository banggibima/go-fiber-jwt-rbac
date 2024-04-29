package jwt

import (
	"time"

	"github.com/banggibima/go-fiber-jwt-rbac/config"
	"github.com/banggibima/go-fiber-jwt-rbac/internal/domain/entity"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(config *config.Config, user *entity.User) (interface{}, error) {
	accessExpiry := config.JWT.AccessExpiry
	accessSecret := config.JWT.AccessSecret

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aud": config.JWT.Audience,
		"exp": time.Now().Add(time.Duration(accessExpiry) * time.Second).Unix(),
		"iat": time.Now().Unix(),
		"iss": config.JWT.Issuer,
		"nbf": time.Now().Unix(),
		"sub": map[string]interface{}{
			"id":   user.ID,
			"role": user.Role,
		},
	})

	accessTokenString, err := accessToken.SignedString([]byte(accessSecret))
	if err != nil {
		return nil, err
	}

	refreshExpiry := config.JWT.RefreshExpiry
	refreshSecret := config.JWT.RefreshSecret

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Duration(refreshExpiry) * time.Second).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"sub": map[string]interface{}{
			"id":   user.ID,
			"role": user.Role,
		},
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(refreshSecret))
	if err != nil {
		return nil, err
	}

	token := JWT{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}

	return token, nil
}
