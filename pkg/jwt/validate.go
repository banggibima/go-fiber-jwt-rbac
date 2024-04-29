package jwt

import (
	"github.com/banggibima/go-fiber-jwt-rbac/config"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(config *config.Config, tokenString string) (*jwt.Token, error) {
	accessSecret := config.JWT.AccessSecret
	refreshSecret := config.JWT.RefreshSecret

	accessToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if accessToken.Valid {
		return accessToken, nil
	}

	refreshToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(refreshSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if !refreshToken.Valid {
		return nil, err
	}

	return refreshToken, nil
}
