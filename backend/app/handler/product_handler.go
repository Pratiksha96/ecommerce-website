package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"ecommerce-website/app/manager"
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
		manager.GetAllProducts(w, email)
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
		response, err := productManager.CreateProduct(product, "admin", email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(response)
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
			manager.UpdateProduct(id, product, w, "admin", email)
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
		manager.DeleteProduct(id, w, "admin", email)
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
