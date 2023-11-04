package config

import (
	"github.com/chalermridp/go-template-rest/internal/app/controller"
	"github.com/chalermridp/go-template-rest/internal/app/repository"
	"github.com/chalermridp/go-template-rest/internal/app/service"
)

type Initialization struct {
	controller.UserController
	service.UserService
	repository.UserRepository
}

func NewInitialization(
	userController controller.UserController,
	userService service.UserService,
	userRepository repository.UserRepository) *Initialization {
	return &Initialization{UserController: userController, UserService: userService, UserRepository: userRepository}
}
