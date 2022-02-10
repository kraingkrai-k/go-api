package bbl

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"go-api/src/model"
)

var merchant model.MerchantConnectBBL

func GenerateSecureHash(input model.DataFeed) (hash []byte, err error) {
	init := merchant.NewMerchantConnectBBL()
	shaHash := sha512.New()

	//fmt.Println("input", input.SecureHash)
	secureHash := fmt.Sprintf("%s|%s|%s|%s|%s|%s",
		init.MerchantId,
		input.RefOrderId,
		input.Currency,
		input.Amount,
		"H",
		init.Hash,
	)

	parseShaHash, err := json.Marshal(secureHash)
	if err != nil {
		return nil, err
	}

	fmt.Println("write with" , string(parseShaHash))
	shaHash.Write(parseShaHash)
	return shaHash.Sum(nil), err
}

func VerifySecureHash(secureHash, dataFeedHash string) bool {
	if secureHash != dataFeedHash {
		return false
	}
	return true
}
