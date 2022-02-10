package model

type MemberPay struct {
	ResponseStatus Response      `xml:"responsestatus" json:"responseStatus"`
	Data           MemberPayData `xml:"response" json:"data"`
}

type MemberPayData struct {
	Accounts []MemberPayAccounts `xml:"accounts" json:"accounts"`
}

type MemberPayAccounts struct {
	AccountID   string `xml:"accountId" json:"accountID"`
	CardType    string `xml:"accounttype" json:"cardType"`
	CardNo      string `xml:"account" json:"cardNo"`
	ExpYear     string `xml:"expyear" json:"expYear"`
	ExpMonth    string `xml:"expmonth" json:"expMonth"`
	HolderName  string `xml:"holdername" json:"holderName"`
	StaticToken string `xml:"statictoken" json:"staticToken"`
	Status      string `xml:"status" json:"status"`
}
