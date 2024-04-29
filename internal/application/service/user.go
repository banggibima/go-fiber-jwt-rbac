package service

import (
	"github.com/banggibima/go-fiber-jwt-rbac/config"
	"github.com/banggibima/go-fiber-jwt-rbac/internal/domain/entity"
	"github.com/banggibima/go-fiber-jwt-rbac/internal/domain/repository"
	"github.com/banggibima/go-fiber-jwt-rbac/pkg/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository repository.UserRepository
	Config         *config.Config
}

func NewUserService(
	userRepository repository.UserRepository,
	config *config.Config,
) *UserService {
	return &UserService{
		UserRepository: userRepository,
		Config:         config,
	}
}

func (s *UserService) ReadAll() ([]*entity.User, error) {
	users, err := s.UserRepository.ReadAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) ReadByID(id uuid.UUID) (*entity.User, error) {
	user, err := s.UserRepository.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) ReadByUsername(username string) (*entity.User, error) {
	user, err := s.UserRepository.ReadByUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Create(user *entity.User) error {
	if err := s.UserRepository.Create(user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) Update(id uuid.UUID, user *entity.User) error {
	if err := s.UserRepository.Update(id, user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) Delete(id uuid.UUID) error {
	if err := s.UserRepository.Delete(id); err != nil {
		return err
	}

	return nil
}

func (s *UserService) Login(username, password string) (interface{}, error) {
	user, err := s.ReadByUsername(username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	token, err := jwt.GenerateToken(s.Config, user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) Register(user *entity.User) (interface{}, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user.Password = string(bytes)
	user.Role = "user"

	if err := s.Create(user); err != nil {
		return "", err
	}

	token, err := jwt.GenerateToken(s.Config, user)
	if err != nil {
		return "", err
	}

	return token, nil
}
