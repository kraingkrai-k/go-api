package circle

import "go-api/src/helpers/shape"

type Circle struct {
	circle shape.Interface
}

func New(i shape.Interface) (service shape.Interface) {
	return &Circle{
		circle: i,
	}
}
