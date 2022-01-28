package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"ecommerce-website/app/handler"

	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/ping", handler.PingHandler())

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Starting server...")
	log.Fatal(srv.ListenAndServe())
}
