package controller

import (
	"net/http"

	"github.com/chalermridp/go-template-rest/internal/app/constant"
	"github.com/chalermridp/go-template-rest/internal/app/domain/dao"
	"github.com/chalermridp/go-template-rest/internal/app/pkg"
	"github.com/chalermridp/go-template-rest/internal/app/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type DefaultUserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *DefaultUserController {
	return &DefaultUserController{userService: userService}
}

func (d DefaultUserController) GetAll(ctx *gin.Context) {
	defer pkg.PanicHandler(ctx)

	users, err := d.userService.GetAll()
	if err != nil {
		pkg.PanicException(constant.UnknownError)
	}

	ctx.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, users))
}

func (d DefaultUserController) GetById(ctx *gin.Context) {
	defer pkg.PanicHandler(ctx)

	// todo get and validate request
	id := "1234"

	user, err := d.userService.GetById(id)
	if err != nil {
		pkg.PanicException(constant.UnknownError)
	}
	ctx.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, user))
}

func (d DefaultUserController) Create(ctx *gin.Context) {
	defer pkg.PanicHandler(ctx)

	// todo get and validate request
	user := dao.User{}

	createdUser, err := d.userService.Create(&user)
	if err != nil {
		pkg.PanicException(constant.UnknownError)
	}
	ctx.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, createdUser))
}
