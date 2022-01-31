package model

import (
	"time"
)

type Paginate struct {
	TotalRecords int32
	Total        int32
}

type Filter struct {
	Limit     int32  `form:"limit" validate:"required"`
	Offset    int32  `form:"offset"`
	SortKey   string `form:"sortKey"`
	SortBy    string `form:"sortBy, oneof=asc,desc"`
	StartDate time.Time
	EndDate   time.Time
}

func (f Filter) Validate() error {
	if err := validate.Struct(f); err != nil {
		return err
	}

	return nil
}
