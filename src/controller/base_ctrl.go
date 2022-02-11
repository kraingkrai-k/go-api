package controller

import "github.com/kraingkrai-k/go-api/src/usecase"

type Controller struct {
	usecase usecase.Interface
}

func New(uc usecase.Interface) *Controller {
	return &Controller{
		usecase: uc,
	}
}
