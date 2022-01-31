package router

import (
	"go-api/src/controller"
	"go-api/src/repository"
	"go-api/src/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRest(app *gin.Engine, dbSQL *gorm.DB) *gin.Engine {

	repo := repository.New(dbSQL)
	uc := usecase.New(repo)
	ctrl := controller.New(uc)

	route := app.Group("/api")
	// baseURL := ""

	// public route
	route.POST("/signin", ctrl.SigninWithUsername)

	// user route
	userGroup := route.Group("/user", AuthorizedJWT())
	{
		userGroup.POST("/", ctrl.AddUser)
		userGroup.GET("/", ctrl.GetUsers)
		userGroup.GET("/:id", ctrl.GetUser)
		userGroup.PUT("/:id", ctrl.UpdateUser)
		userGroup.DELETE("/:id", ctrl.DeleteUser)
	}

	// routes
	return app
}
