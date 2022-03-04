package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"ecommerce-website/app/middleware"
	models "ecommerce-website/app/models"
	"ecommerce-website/app/utils"
	"log"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		ctx := r.Context()
		email := ctx.Value("email").(string)
		middleware.GetAllProducts(w, email)
	}
}

func SearchProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		ctx := r.Context()
		email := ctx.Value("email").(string)
		query := r.URL.Query()
		middleware.SearchProducts(w, query, email)
	}
}

func CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		var product models.Product
		_ = json.NewDecoder(r.Body).Decode(&product)

		if errors := utils.Validate(product); len(errors) > 0 {
			log.Println("Received invalid json request!")
			err := map[string]interface{}{"success": false, "message": errors}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
		} else {
			ctx := r.Context()
			email := ctx.Value("email").(string)
			middleware.CreateProduct(product, w, "admin", email)
		}
	}
}

func UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var params = mux.Vars(r)
		id, _ := primitive.ObjectIDFromHex(params["id"])

		var product models.Product
		_ = json.NewDecoder(r.Body).Decode(&product)

		if errors := utils.Validate(product); len(errors) > 0 {
			log.Println("Received invalid json request!")
			err := map[string]interface{}{"success": false, "message": errors}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
		} else {
			ctx := r.Context()
			email := ctx.Value("email").(string)
			middleware.UpdateProduct(id, product, w, "admin", email)
		}
	}
}

func DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var params = mux.Vars(r)
		if params["id"] == "" {
			utils.GetError(errors.New("input id for delete is invalid"), w)
			return
		}

		id, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil {
			utils.GetError(errors.New("invalid object id"), w)
			return
		}
		ctx := r.Context()
		email := ctx.Value("email").(string)
		middleware.DeleteProduct(id, w, "admin", email)
	}
}

func GetProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var params = mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil {
			utils.GetError(errors.New("invalid object id"), w)
			return
		}
		ctx := r.Context()
		email := ctx.Value("email").(string)
		middleware.GetProduct(id, w, email)
	}
}
