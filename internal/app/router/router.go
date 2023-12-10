package router

import (
	"github.com/chalermridp/go-template-rest/internal/app/config"
	"github.com/gin-gonic/gin"
)

func NewRouter(init *config.Initialization) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		user := api.Group("/users")
		user.GET("", init.UserController.GetAll)
		user.GET("/:userID", init.UserController.GetById)
		user.POST("", init.UserController.Create)
	}
	return router
}
