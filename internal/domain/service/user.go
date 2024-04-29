package service

import (
	"github.com/banggibima/go-fiber-jwt-rbac/internal/domain/entity"
	"github.com/google/uuid"
)

type UserService interface {
	ReadAll() ([]*entity.User, error)
	ReadByID(id uuid.UUID) (*entity.User, error)
	ReadByUsername(username string) (*entity.User, error)
	Create(user *entity.User) error
	Update(id uuid.UUID, user *entity.User) error
	Delete(id uuid.UUID) error
	Login(username, password string) (interface{}, error)
	Register(user *entity.User) (interface{}, error)
}
