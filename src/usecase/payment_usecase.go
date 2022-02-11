package usecase

import (
	"bytes"
	"crypto/sha512"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kraingkrai-k/go-api/src/helpers/bbl"
	bblModel "github.com/kraingkrai-k/go-api/src/helpers/bbl/model"
	"github.com/kraingkrai-k/go-api/src/model"
	"io/ioutil"
	"log"
	"net/http"
)

var url = "https://psipay.bangkokbank.com/b2c/eng/"

func (uc *Usecase) RegisterBBLCard(input model.RegisterCard) (output model.RegisterCard, err error) {

	pathAddMember := fmt.Sprintf("%smerchant/go-api/MemberPayApi.jsp", url)
	bodyString := fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s", input.MerchantId, input.MerchantApiId, input.Password, input, input.Amount, input.PaymentType, input.Hash)
	shaHash := sha512.New()

	parseShaHash, err := json.Marshal(bodyString)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	shaHash.Write(parseShaHash)
	fmt.Printf("sha512: %x\n", shaHash.Sum(nil))

	log.Printf("bodyString\n", bodyString)
	bBody, err := json.Marshal(input)
	if err != nil {
		return output, err
	}

	responseBody := bytes.NewBuffer(bBody)

	resp, err := http.NewRequest("POST", pathAddMember, responseBody)
	if err != nil {
		return output, err
	}

	defer resp.Body.Close()
	parse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return output, err
	}
	sb := string(parse)
	log.Printf(sb)

	return output, nil
}

func (uc *Usecase) CallbackDataFeed(input model.DataFeed) (output model.DataFeed, err error) {

	secureHash, err := bbl.GenerateSecureHash(input)
	if err != nil {
		return input, err
	}

	x := string(secureHash)
	fmt.Println("from generate", x)

	ok := bbl.VerifySecureHash(x, input.SecureHash)
	if !ok {
		return input, nil
	}

	//fmt.Println("success out", string(secureHash))
	return input, nil
}

func (uc *Usecase) GetMemberPay(input bblModel.BBLRequest) (output []bblModel.MemberPayAccounts, err error) {

	result, err := uc.bbl.QueryMemberPay(input)
	if err != nil {
		return output, err
	}

	if result.ResponseStatus.ResponseCode != "0" {
		return output, errors.New(result.ResponseStatus.ResponseMessage)
	}

	if err != nil {
		return output, err
	}

	output = result.Data.Accounts
	return output, nil
}
