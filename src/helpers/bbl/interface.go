package bbl

import "go-api/src/helpers/bbl/model"

type Service interface {
	QueryMemberPay(input model.BBLRequest) (output model.MemberPay, err error)
}
