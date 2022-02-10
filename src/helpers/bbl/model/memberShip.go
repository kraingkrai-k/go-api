package model

type MemberShip struct {
	ResponseStatus Response      `xml:"responsestatus" json:"responseStatus"`
	Data           MemberShipData `xml:"response" json:"data"`
}

type MemberShipData struct {
	Accounts []MemberShipAccounts `xml:"accounts" json:"accounts"`
}

type MemberShipAccounts struct {
	AccountID   string `xml:"accountId" json:"accountID"`
	CardType    string `xml:"accounttype" json:"cardType"`
	CardNo      string `xml:"account" json:"cardNo"`
	ExpYear     string `xml:"expyear" json:"expYear"`
	ExpMonth    string `xml:"expmonth" json:"expMonth"`
	HolderName  string `xml:"holdername" json:"holderName"`
	StaticToken string `xml:"statictoken" json:"staticToken"`
	Status      string `xml:"status" json:"status"`
}
