package controller

import "go-api/src/usecase"

type Controller struct {
	usecase usecase.Interface
}

func New(uc usecase.Interface) *Controller {
	return &Controller{
		usecase: uc,
	}
}
