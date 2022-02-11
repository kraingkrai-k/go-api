package router

import (
	"github.com/kraingkrai-k/go-api/src/controller"
	bblImpl "github.com/kraingkrai-k/go-api/src/helpers/bbl/implement"
	bblModel "github.com/kraingkrai-k/go-api/src/helpers/bbl/model"
	"github.com/kraingkrai-k/go-api/src/helpers/shape/implement/circle"
	"github.com/kraingkrai-k/go-api/src/helpers/shape/implement/rectangle"
	"github.com/kraingkrai-k/go-api/src/helpers/wraphttps"
	"github.com/kraingkrai-k/go-api/src/repository"
	"github.com/kraingkrai-k/go-api/src/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRest(app *gin.Engine, dbSQL *gorm.DB) *gin.Engine {

	repo := repository.New(dbSQL)

	shapeRecImpl := rectangle.New(&rectangle.Rectangle{})
	shapeCircleImpl := circle.New(&circle.Circle{})

	bbl := bblImpl.New(wraphttps.New(), bblModel.NewBBL())
	uc := usecase.New(repo, shapeRecImpl, shapeCircleImpl, bbl)
	ctrl := controller.New(uc)

	route := app.Group("/api")
	// baseURL := ""

	// public route
	route.POST("/signin", ctrl.SigninWithUsername)

	// Dev
	route.POST("/feed", ctrl.DataFeed)

	route.POST("/register/bbl/card", ctrl.RegisterBBLCard)
	route.POST("/queryMember/bbl", ctrl.GetMemberPay)

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
