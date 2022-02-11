package implement

import (
	"github.com/kraingkrai-k/go-api/src/helpers/bbl"
	"github.com/kraingkrai-k/go-api/src/helpers/bbl/model"
	"github.com/kraingkrai-k/go-api/src/helpers/wraphttps"
)

type implement struct {
	http wraphttps.WrapHttps
	info *model.BBL
}

type Action string

func (a Action) ToString() string {
	return string(a)
}

const (
	QueryAccount Action = "QueryAccount"
)

const (
	MemberPayUrl  = "MemberPayApi.jsp"
	MemberShipUrl = "MembershipApi.jsp"
)

func New(client wraphttps.WrapHttps, info *model.BBL) (service bbl.Service) {
	return &implement{
		http: client,
		info: info,
	}
}
