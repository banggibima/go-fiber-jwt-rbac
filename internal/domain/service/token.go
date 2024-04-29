package service

import (
	"github.com/banggibima/go-fiber-jwt-rbac/internal/domain/entity"
)

type TokenService interface {
	ReadByRefreshToken(refreshToken string) (*entity.Token, error)
	Create(token *entity.Token) error
	DeleteByRefreshToken(refreshToken string) error
}
