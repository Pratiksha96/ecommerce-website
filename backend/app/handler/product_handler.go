package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"ecommerce-website/app/manager"
	models "ecommerce-website/app/models"
	"ecommerce-website/app/utils"
	"log"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllProducts(productManager manager.ProductManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		ctx := r.Context()
		email := ctx.Value("email").(string)
		payload, err := productManager.GetAllProducts(email)
		if err != nil {
			utils.GetErrorWithStatus(errors.New("invalid object id"), w, http.StatusUnprocessableEntity)
			return
		}
		json.NewEncoder(w).Encode(payload)
	}
}

func SearchProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		ctx := r.Context()
		email := ctx.Value("email").(string)
		query := r.URL.Query()
		manager.SearchProducts(w, query, email)
	}
}

func CreateProduct(productManager manager.ProductManager) http.HandlerFunc {
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
			return
		}
		ctx := r.Context()
		email := ctx.Value("email").(string)
		fmt.Println("Creating product")
		response, err := productManager.CreateProduct(product, "admin", email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

func UpdateProduct(productManager manager.ProductManager) http.HandlerFunc {
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
			product, err := productManager.UpdateProduct(id, product, "admin", email)
			if err != nil {
				utils.GetError(err, w)
				return
			}
			json.NewEncoder(w).Encode(product)
		}
	}
}

func DeleteProduct(productManager manager.ProductManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var params = mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil || params["id"] == "" {
			utils.GetErrorWithStatus(errors.New("invalid object id"), w, http.StatusUnprocessableEntity)
			return
		}
		ctx := r.Context()
		email := ctx.Value("email").(string)
		deleteResponse, err := productManager.DeleteProduct(id, "admin", email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(deleteResponse)
	}
}

func GetProduct(productManager manager.ProductManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var params = mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil {
			utils.GetErrorWithStatus(errors.New("invalid object id"), w, http.StatusUnprocessableEntity)
			return
		}
		ctx := r.Context()
		email := ctx.Value("email").(string)
		product, err := productManager.GetProduct(id, email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(product)
	}
}
