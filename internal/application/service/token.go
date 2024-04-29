package service

import (
	"github.com/banggibima/go-fiber-jwt-rbac/config"
	"github.com/banggibima/go-fiber-jwt-rbac/internal/domain/entity"
	"github.com/banggibima/go-fiber-jwt-rbac/internal/domain/memory"
)

type TokenService struct {
	TokenMemory memory.TokenMemory
	Config      *config.Config
}

func NewTokenService(
	userMemory memory.TokenMemory,
	config *config.Config,
) *TokenService {
	return &TokenService{
		TokenMemory: userMemory,
		Config:      config,
	}
}

func (s *TokenService) ReadByRefreshToken(refreshToken string) (*entity.Token, error) {
	token, err := s.TokenMemory.ReadByRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *TokenService) Create(token *entity.Token) error {
	if err := s.TokenMemory.Create(token); err != nil {
		return err
	}

	return nil
}

func (s *TokenService) DeleteByRefreshToken(refreshToken string) error {
	if err := s.TokenMemory.DeleteByRefreshToken(refreshToken); err != nil {
		return err
	}

	return nil
}
