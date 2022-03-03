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

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	isValid, email := middleware.AuthenticateUser(w, r)
	if isValid {
		middleware.GetAllProducts(w, email)
	}
}

func SearchProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	query := r.URL.Query()
	isValid, email := middleware.AuthenticateUser(w, r)
	if isValid {
		middleware.SearchProducts(w, query, email)
	}

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
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
		isValid, email := middleware.AuthenticateUser(w, r)
		if isValid {
			middleware.CreateProduct(product, w, "admin", email)
		}
	}

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
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
		isValid, email := middleware.AuthenticateUser(w, r)
		if isValid {
			middleware.UpdateProduct(id, product, w, "admin", email)
		}
	}
}

//delete product using product id
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// get params
	var params = mux.Vars(r)

	if params["id"] == "" {
		utils.GetError(errors.New("input id for delete is invalid"), w)
		return
	}
	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		utils.GetError(errors.New("invalid object id"), w)
		return
	}
	isValid, email := middleware.AuthenticateUser(w, r)
	if isValid {
		middleware.DeleteProduct(id, w, "admin", email)
	}
}

//get product details by sending a product id
func GetProduct(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	// we get params with mux.
	var params = mux.Vars(r)
	// string to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		utils.GetError(errors.New("invalid object id"), w)
		return
	}
	isValid, email := middleware.AuthenticateUser(w, r)
	if isValid {
		middleware.GetProduct(id, w, email)
	}
}
