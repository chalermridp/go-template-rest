//go:build wireinject
// +build wireinject

package config

import (
	"github.com/chalermridp/go-template-rest/internal/app/controller"
	"github.com/chalermridp/go-template-rest/internal/app/repository"
	"github.com/chalermridp/go-template-rest/internal/app/service"
	"github.com/google/wire"
)

var userRepoSet = wire.NewSet(repository.NewUserRepository, wire.Bind(new(repository.UserRepository), new(*repository.DefaultUserRepository)))
var userServiceSet = wire.NewSet(service.NewUserService, wire.Bind(new(service.UserService), new(*service.DefaultUserService)))
var userControllerSet = wire.NewSet(controller.NewUserController, wire.Bind(new(controller.UserController), new(*controller.DefaultUserController)))

func Init() *Initialization {
	wire.Build(NewInitialization, userRepoSet, userServiceSet, userControllerSet)
	return nil
}
