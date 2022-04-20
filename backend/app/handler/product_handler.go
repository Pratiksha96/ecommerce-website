package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"ecommerce-website/app/manager"
	models "ecommerce-website/app/models"
	"ecommerce-website/app/utils"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllProducts(productManager manager.ProductManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		payload, err := productManager.GetAllProducts()
		if err != nil {
			utils.GetErrorWithStatus(errors.New("invalid object id"), w, http.StatusUnprocessableEntity)
			return
		}
		json.NewEncoder(w).Encode(payload)
	}
}

func SearchProducts(productManager manager.ProductManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		query := r.URL.Query()
		payload, err := productManager.SearchProducts(query)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(payload)
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

		if err := utils.ProductValidation(product); len(err) > 0 {
			utils.GetErrorWithStatus(errors.New("Received invalid json request!"), w, http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		email := ctx.Value("email").(string)
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

		if err := utils.ProductValidation(product); len(err) > 0 {
			utils.GetErrorWithStatus(errors.New("Received invalid json request!"), w, http.StatusBadRequest)
			return
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
