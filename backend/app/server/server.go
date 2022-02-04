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
	r.HandleFunc("/product/get", handler.GetAllProducts).Methods("GET", "OPTIONS")
	r.HandleFunc("/product/add", handler.CreateProduct).Methods("POST", "OPTIONS")
	r.HandleFunc("/product/update/{id}", handler.UpdateProduct).Methods("PUT", "OPTIONS")
	r.HandleFunc("/product/delete/{id}", handler.DeleteProduct).Methods("DELETE", "OPTIONS")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Starting server...")
	log.Fatal(srv.ListenAndServe())
}
