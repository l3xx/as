package main

import (
	"log"
	"net/http"
	"travel/internal/pkg/config"
	"travel/pkg/middleware"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		panic("Not load config")
	}

	mw := middleware.NewMiddleware(cfg)
	err = http.ListenAndServe(":8080", mw)
	if err != nil {
		log.Fatal("Error start ", err)
	}
}
