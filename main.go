package main

import (
	"github.com/kraingkrai-k/go-api/src/config"
	"github.com/kraingkrai-k/go-api/src/router"
	"net/http"
	"time"
)

func main() {

	_ = config.New()
	dbSQL := config.NewSQL()
	route := router.NewRouter()
	app := router.NewRest(route, dbSQL)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      app,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	_ = s.ListenAndServe()
}

//type contentType string
//
//const (
//	XML  contentType = "XML"
//	JSON contentType = "JSON"
//)
//
//func (t contentType) toString() string {
//	return string(t)
//}
//
//type X struct {
//	BaseUri string `json:"baseUri"`
//	Title   string `json:"title,omitempty"`
//}
//
//type Y struct {
//	Title string `json:"title,omitempty"`
//}
//
//func (x X) MakeUrl(ct contentType) string {
//	switch ct {
//	case JSON:
//		return x.BaseUri + "json"
//	case XML:
//		return x.BaseUri + "xml"
//	default:
//		return ""
//	}
//}
//
//func (x X) Info() Y {
//	fmt.Println("Info" , x)
//	return Y{
//		Title: x.Title,
//	}
//}
//
//type Interface interface {
//	MakeUrl(ct contentType) string
//	Info() Y
//}
//
//func NewX() Interface {
//	return &X{
//		Title:   "Title",
//		BaseUri: "foobar/go-api/",
//	}
//}
//
//func main() {
//	a := NewX()
//
//	//var b = X{}
//	b := X{}
//
//	fmt.Println("log b before", b.Info())
//	fmt.Println("log a", a)
//	b.Info()
//	fmt.Println("log b after", b)
//	marshal, err := json.Marshal(a)
//	if err != nil {
//		return
//	}
//
//	fmt.Println("ms", string(marshal))
//}
