package bbl

import "github.com/kraingkrai-k/go-api/src/helpers/bbl/model"

type Service interface {
	QueryMemberPay(input model.BBLRequest) (output model.MemberPay, err error)
}
