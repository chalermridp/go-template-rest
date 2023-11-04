package service

import (
	"github.com/chalermridp/go-template-rest/internal/app/domain/dao"
	"github.com/chalermridp/go-template-rest/internal/app/repository"
)

type UserService interface {
	GetAll() ([]*dao.User, error)
	GetById(id string) (*dao.User, error)
	Create(user *dao.User) (*dao.User, error)
}

func NewUserService(userRepository repository.UserRepository) *DefaultUserService {
	return &DefaultUserService{userRepository: userRepository}
}

type DefaultUserService struct {
	userRepository repository.UserRepository
}

func (d DefaultUserService) GetAll() ([]*dao.User, error) {
	return d.userRepository.GetAll()
}

func (d DefaultUserService) GetById(id string) (*dao.User, error) {
	return d.userRepository.GetById(id)
}

func (d DefaultUserService) Create(user *dao.User) (*dao.User, error) {
	return d.Create(user)
}
