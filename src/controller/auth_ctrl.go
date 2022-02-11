package controller

import (
	"github.com/kraingkrai-k/go-api/src/helpers/response"
	"github.com/kraingkrai-k/go-api/src/middleware"
	"github.com/kraingkrai-k/go-api/src/model"

	"github.com/gin-gonic/gin"
)

var middlewareJWT = middleware.NewJWT()

func (c *Controller) SigninWithUsername(ctx *gin.Context) {

	var bodySignin model.SigninWithUsername

	if err := ctx.ShouldBindJSON(&bodySignin); err != nil {
		response.Error(ctx, err)
		return
	}

	if err := bodySignin.Validate(); err != nil {
		response.Error(ctx, err)
		return
	}

	token := middlewareJWT.GenerateToken(bodySignin.Username, bodySignin.Password)

	response.Success(ctx, token)
}
