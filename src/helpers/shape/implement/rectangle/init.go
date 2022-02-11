package rectangle

import "github.com/kraingkrai-k/go-api/src/helpers/shape"

type Rectangle struct {
	rectangle shape.Interface
}

func New(i shape.Interface) (service shape.Interface) {
	return &Rectangle{
		rectangle: i,
	}
}
