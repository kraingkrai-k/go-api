package model

type DataFeed struct {
	StatusCodePrimary            string `json:"statusCodePrimary" form:"src"`
	StatusCode                   string `json:"statusCode" form:"prc"`
	BankRefOrderId               string `json:"bankRefOrderId" form:"Ord"`
	Holder                       string `json:"nameHolder" form:"Holder"`
	TransactionStatus            string `json:"transactionStatus" form:"successcode"`
	RefOrderId                   string `json:"refOrderId" form:"Ref"`
	OrderRef1                    string `json:"orderRef1" form:"orderRef1"`
	OrderRef2                    string `json:"orderRef2" form:"orderRef2"`
	OrderRef3                    string `json:"orderRef3" form:"orderRef3"`
	RefPaymentNumber             string `json:"refPaymentNumber" form:"PayRef"`
	Amount                       string `json:"amount" form:"Amt"`
	PreDigit                     string `json:"preDigit" form:"cc0104"`
	LastDigit                    string `json:"lastDigit" form:"cc1316"`
	Currency                     string `json:"currency" form:"Cur"`
	Remark                       string `json:"remark" form:"remark"`
	ApproveCode                  string `json:"approveCode" form:"authId"`
	Eci                          string `json:"eci" form:"eci"`
	PayerAuth                    string `json:"payerAuth" form:"payerAuth"`
	SourceIp                     string `json:"sourceIp" form:"sourceIp"`
	TransactionTime              string `json:"transactionTime" form:"TxTime"`
	SecureHash                   string `json:"secureHash" form:"secureHash"`
	PaymentMemberId              string `json:"mpMemberId" form:"mpMemberId"`
	PaymentMemberLastStaticToken string `json:"mpLatestStaticToken" form:"mpLatestStaticToken"`
}

const (
	ADD = "Add"
)

func ActionTypeToString(act string) string {
	if act != "" {
		return ADD
	}
	return act
}

type MerchantConnectBBL struct {
	MerchantId    string `json:"merchantId"`
	MerchantApiId string `json:"merchantApiId"`
	Password      string `json:"password"`
	Hash          string `json:"hash"`
	CurrencyCode  string `json:"currencyCode"`
}

func (m *MerchantConnectBBL) NewMerchantConnectBBL() *MerchantConnectBBL {
	return &MerchantConnectBBL{
		MerchantId:    "9915",
		MerchantApiId: "pwc01",
		Password:      "Api9915",
		Hash:          "PmmMqQaB3doJUhUAM2NC8ih9RNCYUjKj",
		CurrencyCode:  "764",
	}
}

type RegisterCard struct {
	MerchantConnectBBL
	ActionType  string `json:"actionType"`
	MemberId    string `json:"memberId"`
	Account     string `json:"account"`
	ExpMonth    string `json:"expMonth"`
	ExpYear     string `json:"expYear"`
	HolderName  string `json:"holderName"`
	Email       string `json:"email"`
	Amount      string `json:"amount"`
	PaymentType string `json:"paymentType"`
}
