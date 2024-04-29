package memory

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/banggibima/go-fiber-jwt-rbac/internal/domain/entity"
	"github.com/redis/go-redis/v9"
)

type TokenMemory struct {
	RDB   *redis.Client
	Mutex *sync.RWMutex
}

func NewTokenMemory(
	rdb *redis.Client,
) *TokenMemory {
	return &TokenMemory{
		RDB:   rdb,
		Mutex: &sync.RWMutex{},
	}
}

func (m *TokenMemory) ReadByRefreshToken(refreshToken string) (*entity.Token, error) {
	m.Mutex.RLock()
	defer m.Mutex.RUnlock()

	ctx := context.Background()
	tokenBytes, err := m.RDB.Get(ctx, refreshToken).Bytes()
	if err != nil {
		return nil, err
	}

	token := new(entity.Token)
	if err := json.Unmarshal(tokenBytes, token); err != nil {
		return nil, err
	}

	return token, nil
}

func (m *TokenMemory) Create(token *entity.Token) error {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	ctx := context.Background()
	tokenBytes, err := json.Marshal(token)
	if err != nil {
		return err
	}

	if _, err := m.RDB.Set(ctx, token.RefreshToken, tokenBytes, 0).Result(); err != nil {
		return err
	}

	return nil
}

func (m *TokenMemory) DeleteByRefreshToken(refreshToken string) error {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	ctx := context.Background()
	if _, err := m.RDB.Del(ctx, refreshToken).Result(); err != nil {
		return err
	}

	return nil
}
