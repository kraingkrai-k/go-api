package controller

import (
	"fmt"
	bblModel "github.com/kraingkrai-k/go-api/src/helpers/bbl/model"
	"github.com/kraingkrai-k/go-api/src/helpers/response"
	"github.com/kraingkrai-k/go-api/src/model"

	"github.com/gin-gonic/gin"
)

func (c *Controller) DataFeed(ctx *gin.Context) {

	var feed model.DataFeed

	//err := ctx.ShouldBindQuery(&feed)
	//if err != nil {
	//	response.Error(ctx, err)
	//	return
	//}

	err := ctx.ShouldBind(&feed)
	if err != nil {
		fmt.Println("err" , err)
		response.Error(ctx, err)
		return
	}

	fmt.Println("log" , feed)
	output, err := c.usecase.CallbackDataFeed(feed)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, output)
}

func (c *Controller) GetMemberPay(ctx *gin.Context) {
	var bblReq bblModel.BBLRequest

	err := ctx.ShouldBindJSON(&bblReq)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	output, err := c.usecase.GetMemberPay(bblReq)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, output)
}

func (c *Controller) RegisterBBLCard(ctx *gin.Context) {

	var register model.RegisterCard

	err := ctx.ShouldBindJSON(&register)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	result, err := c.usecase.RegisterBBLCard(register)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, result)
}
