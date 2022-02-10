package wraphttps

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	ContentTypeStr             = "Content-Type"
	Get            Method      = "GET"
	Post           Method      = "POST"
	Put            Method      = "PUT"
	Delete         Method      = "DELETE"
	Json           ContentType = "application/json"
	FormData       ContentType = "multipart/form-data"
	XML            ContentType = "application/xml"
	URLEncoded     ContentType = "application/x-www-form-urlencoded"
)

type Method string

func (m Method) ToString() string {
	return string(m)
}

type ContentType string

func (c ContentType) ToString() string {
	return string(c)
}

type WrapHttps struct {
	http.Client
}

func New() WrapHttps {
	return WrapHttps{
		http.Client{},
	}
}

func (w WrapHttps) MakeRequest(method Method, uri string, contentType ContentType, iBody interface{}) (resp []byte, err error) {

	body, err := MakeBody(contentType, iBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method.ToString(), uri, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(ContentTypeStr, contentType.ToString())
	res, err := w.Do(req)
	if err != nil {
		return
	}
	defer func() { _ = res.Body.Close() }()

	return ioutil.ReadAll(res.Body)
}

func MakeBody(contentType ContentType, iBody interface{}) (resp io.Reader, err error) {
	switch contentType {
	case URLEncoded:
		body, ok := iBody.(map[string]string)
		if !ok {
			return nil, errors.New("invalid parse")
		}

		data := url.Values{}
		for k, v := range body {
			data.Set(k, v)
		}
		return strings.NewReader(data.Encode()), nil
	default:
		return nil, errors.New("not yet")
	}
}
