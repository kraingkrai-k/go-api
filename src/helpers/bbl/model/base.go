package model

import "go-api/src/config"

type BBL struct {
	BaseBBLURL    string
	MerchantId    string `json:"merchantId"`
	MerchantApiId string `json:"merchantApiId"`
	Password      string `json:"password"`
	Hash          string `json:"hash"`
	CurrencyCode  string `json:"currencyCode"`
}

func NewBBL() *BBL {
	x := config.New()
	return &BBL{
		BaseBBLURL:    x.GetBBLUrl(),
		MerchantId:    "9915",
		MerchantApiId: "ffapi",
		Password:      "Api9915",
		Hash:          "PmmMqQaB3doJUhUAM2NC8ih9RNCYUjKj",
		CurrencyCode:  "764",
	}
}

type BBLRequest struct {
	Amount     string `json:"amount,omitempty"`
	OrderRef   string `json:"orderRef,omitempty"`
	ActionType string `json:"actionType"`
	MemberId   string `json:"memberId,omitempty"`
	*BBL
}

type Response struct {
	ResponseCode    string `xml:"responsecode" json:"code"`
	ResponseMessage string `xml:"responsemessage" json:"message"`
}
