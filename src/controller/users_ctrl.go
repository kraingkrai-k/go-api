package controller

import (
	"go-api/src/helpers/response"
	"go-api/src/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetUsers(ctx *gin.Context) {
	var filter model.Filter

	if err := ctx.ShouldBindQuery(&filter); err != nil {
		response.Error(ctx, err)
		return
	}

	if err := filter.Validate(); err != nil {
		response.Error(ctx, err)
		return
	}

	if err := model.ValidateUsersSortKey(filter.SortKey); err != nil {
		response.Error(ctx, err)
		return
	}

	users, paginate, err := c.usecase.GetUsers(filter)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, users, paginate)
}

func (c *Controller) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	userId, err := strconv.Atoi(id)
	if err != nil {
		// custom
		// response.Error(ctx, response.Response{
		// 	Status: http.StatusBadGateway,
		// 	Error:  err.Error(),
		// })

		response.Error(ctx, err)
		return
	}

	users, err := c.usecase.GetUser(int32(userId))
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, users)
}

func (c *Controller) AddUser(ctx *gin.Context) {
	input := model.Users{}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.Error(ctx, err)
		return
	}

	if err := input.Validate(); err != nil {
		response.Error(ctx, err)
		return
	}

	users, err := c.usecase.AddUser(input)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, users)
}

func (c *Controller) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		response.Error(ctx, err)
		return
	}

	input := model.Users{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.Error(ctx, err)
		return
	}

	if err := input.Validate(); err != nil {
		response.Error(ctx, err)
		return
	}

	users, err := c.usecase.UpdateUser(input, int32(id))
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, users)
}

func (c *Controller) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		response.Error(ctx, err)
		return
	}

	err = c.usecase.DeleteUser(int32(id))
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, nil)
}
