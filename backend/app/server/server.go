package server

import (
	"log"
	"net/http"
	"time"

	"ecommerce-website/app/handler"
	"ecommerce-website/app/manager"
	"ecommerce-website/app/middleware"

	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()

	productManager := manager.NewProductManager()

	r.HandleFunc("/ping", handler.PingHandler())
	r.HandleFunc("/product/get", middleware.AuthenticateUser(handler.GetAllProducts())).Methods("GET", "OPTIONS")
	r.HandleFunc("/product/search", middleware.AuthenticateUser(handler.SearchProducts())).Methods("GET", "OPTIONS")
	r.HandleFunc("/product/add", middleware.AuthenticateUser(handler.CreateProduct(productManager))).Methods("POST", "OPTIONS")
	r.HandleFunc("/product/update/{id}", middleware.AuthenticateUser(handler.UpdateProduct())).Methods("PUT", "OPTIONS")
	r.HandleFunc("/product/delete/{id}", middleware.AuthenticateUser(handler.DeleteProduct())).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/product/get/{id}", middleware.AuthenticateUser(handler.GetProduct(productManager))).Methods("GET", "OPTIONS")

	r.HandleFunc("/register", handler.RegisterUser()).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", handler.LoginUser()).Methods("POST", "OPTIONS")
	r.HandleFunc("/logout", handler.LogoutUser()).Methods("POST", "OPTIONS")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Starting server...")
	log.Fatal(srv.ListenAndServe())
}
