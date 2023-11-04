package repository

import (
	"errors"

	"github.com/chalermridp/go-template-rest/internal/app/domain/dao"
)

type UserRepository interface {
	GetAll() ([]*dao.User, error)
	GetById(id string) (*dao.User, error)
	Create(user *dao.User) (*dao.User, error)
}

func NewUserRepository() *DefaultUserRepository {
	return &DefaultUserRepository{}
}

type DefaultUserRepository struct {
}

func (u DefaultUserRepository) GetAll() ([]*dao.User, error) {
	//return nil, errors.New("unimplemented")
	return []*dao.User{}, nil
}

func (u DefaultUserRepository) GetById(id string) (*dao.User, error) {
	return nil, errors.New("unimplemented")
}

func (u DefaultUserRepository) Create(user *dao.User) (*dao.User, error) {
	return nil, errors.New("unimplemented")
}
