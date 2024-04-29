package repository

import (
	"sync"

	"github.com/banggibima/go-fiber-jwt-rbac/internal/domain/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB    *gorm.DB
	Mutex *sync.Mutex
}

func NewUserRepository(
	db *gorm.DB,
) *UserRepository {
	return &UserRepository{
		DB:    db,
		Mutex: &sync.Mutex{},
	}
}

func (r *UserRepository) ReadAll() ([]*entity.User, error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	users := make([]*entity.User, 0)

	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) ReadByID(id uuid.UUID) (*entity.User, error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	user := new(entity.User)

	if err := r.DB.Where("id = ?", id).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) ReadByUsername(username string) (*entity.User, error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	user := new(entity.User)

	if err := r.DB.Where("username = ?", username).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Create(user *entity.User) error {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	if err := r.DB.Create(user).Error; err != nil {
		return err
	}

	if err := r.DB.First(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(id uuid.UUID, user *entity.User) error {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	if err := r.DB.Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}

	if err := r.DB.First(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(id uuid.UUID) error {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	user := new(entity.User)

	if err := r.DB.Where("id = ?", id).Delete(user).Error; err != nil {
		return err
	}

	return nil
}
