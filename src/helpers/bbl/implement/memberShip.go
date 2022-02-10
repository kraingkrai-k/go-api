package implement

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"go-api/src/helpers/bbl/model"
	"go-api/src/helpers/wraphttps"
)

func (i *implement) QueryMemberShip(input model.BBLRequest) (output model.MemberPay, err error) {

	url := fmt.Sprintf("%s/%s", i.info.BaseBBLURL, MemberShipUrl)
	body := model.BBLRequest{
		MemberId:   input.MemberId,
		ActionType: QueryAccount.ToString(),
		BBL:        i.info,
	}

	mX, err := json.Marshal(body)
	if err != nil {
		return output, err
	}

	mapBody := map[string]string{}
	err = json.Unmarshal(mX, &mapBody)
	if err != nil {
		return output, err
	}

	result, err := i.http.MakeRequest(wraphttps.Post, url, wraphttps.URLEncoded, mapBody)
	if err != nil {
		return output, err
	}

	err = xml.Unmarshal(result, &output)
	if err != nil {
		return output, err
	}

	return output, nil
}
