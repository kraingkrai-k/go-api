package main

import (
	"go-api/src/config"
	"go-api/src/router"
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

	s.ListenAndServe()
}
