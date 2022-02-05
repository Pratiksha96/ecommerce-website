package handler

import (
	"encoding/json"
	"net/http"

	models "ecommerce-website/app/Models"
	"ecommerce-website/app/middleware"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	middleware.GetAllProducts(w)

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	middleware.CreateProduct(product, w)

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	var product models.Product

	_ = json.NewDecoder(r.Body).Decode(&product)

	middleware.UpdateProduct(id, product, w)
}

//delete product using product id
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// get params
	var params = mux.Vars(r)

	// string to primitve.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	middleware.DeleteProduct(id, w)
}

//get product details by sending a product id
func GetProduct(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	middleware.GetProduct(id, w)
}
