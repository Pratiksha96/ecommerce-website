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
	orderManager := manager.NewOrderManager()
	userManager := manager.NewUserManager()
	reviewManager := manager.NewReviewManager()

	r.HandleFunc("/ping", handler.PingHandler())
	r.HandleFunc("/product/get", handler.GetAllProducts(productManager)).Methods("GET", "OPTIONS")
	r.HandleFunc("/product/search", handler.SearchProducts(productManager)).Methods("GET", "OPTIONS")
	r.HandleFunc("/product/add", middleware.AuthenticateUser(handler.CreateProduct(productManager))).Methods("POST", "OPTIONS")
	r.HandleFunc("/product/update/{id}", middleware.AuthenticateUser(handler.UpdateProduct(productManager))).Methods("PUT", "OPTIONS")
	r.HandleFunc("/product/delete/{id}", middleware.AuthenticateUser(handler.DeleteProduct(productManager))).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/product/get/{id}", middleware.AuthenticateUser(handler.GetProduct(productManager))).Methods("GET", "OPTIONS")

	r.HandleFunc("/product/createReview", middleware.AuthenticateUser(handler.CreateReview(reviewManager))).Methods("PUT", "OPTIONS")
	r.HandleFunc("/product/getreviews/{id}", handler.GetProductReviews(reviewManager)).Methods("GET", "OPTIONS")
	r.HandleFunc("/product/updateReview", middleware.AuthenticateUser(handler.UpdateReview(reviewManager))).Methods("PUT", "OPTIONS")
	r.HandleFunc("/product/deleteReview/{id}", middleware.AuthenticateUser(handler.DeleteReview(reviewManager))).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/register", handler.RegisterUser(userManager)).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", handler.LoginUser(userManager)).Methods("POST", "OPTIONS")
	r.HandleFunc("/logout", handler.LogoutUser(userManager)).Methods("POST", "OPTIONS")
	r.HandleFunc("/me", middleware.AuthenticateUser(handler.GetUserDetails(userManager))).Methods("GET", "OPTIONS")
	r.HandleFunc("/password/update", middleware.AuthenticateUser(handler.UpdatePassword(userManager))).Methods("PUT", "OPTIONS")
	r.HandleFunc("/me/update", middleware.AuthenticateUser(handler.UpdateProfile(userManager))).Methods("PUT", "OPTIONS")
	r.HandleFunc("/user/get", middleware.AuthenticateUser(handler.GetAllUsers(userManager))).Methods("GET", "OPTIONS")
	r.HandleFunc("/user/get/{id}", middleware.AuthenticateUser(handler.GetUser(userManager))).Methods("GET", "OPTIONS")
	r.HandleFunc("/user/updaterole", middleware.AuthenticateUser(handler.UpdateRole(userManager))).Methods("PUT", "OPTIONS")
	r.HandleFunc("/user/delete/{id}", middleware.AuthenticateUser(handler.DeleteUser(userManager))).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/order/create", middleware.AuthenticateUser(handler.CreateOrder(orderManager))).Methods("POST", "OPTIONS")
	r.HandleFunc("/order/user/get", middleware.AuthenticateUser(handler.GetUserOrders(orderManager))).Methods("GET", "OPTIONS")
	r.HandleFunc("/order/get/{id}", middleware.AuthenticateUser(handler.GetSingleOrder(orderManager))).Methods("GET", "OPTIONS")
	r.HandleFunc("/order/get", middleware.AuthenticateUser(handler.GetAllOrders(orderManager))).Methods("GET", "OPTIONS")
	r.HandleFunc("/order/delete/{id}", middleware.AuthenticateUser(handler.DeleteOrder(orderManager))).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/order/update/{id}", middleware.AuthenticateUser(handler.UpdateOrder(orderManager))).Methods("PUT", "OPTIONS")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Starting server...")
	log.Fatal(srv.ListenAndServe())
}
