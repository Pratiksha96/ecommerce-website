package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	models "ecommerce-website/app/Models"
	"ecommerce-website/app/utils"

	"ecommerce-website/internal/database"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	//fmt.Println("this is produc + ", product)
	insertResult, err := database.Coll_product.InsertOne(context.TODO(), product)

	if err != nil {
		utils.GetError(err, w)
		return
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)

	// fmt.Println(task, r.Body)

	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	var product models.Product

	filter := bson.M{"_id": id}

	_ = json.NewDecoder(r.Body).Decode(&product)

	var oldProduct models.Product

	// prepare update model.
	update := bson.D{
		{"$set", bson.D{
			{"name", product.Name},
			{"description", product.Description},
			{"price", product.Price},
			{"ratings", product.Ratings},
			{"images", product.Images},
			{"category", product.Category},
			{"Stock", product.Stock},
			{"reviews", product.Reviews},
		}},
	}

	err := database.Coll_product.FindOneAndUpdate(context.TODO(), filter, update).Decode(&oldProduct)

	if err != nil {
		utils.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// get params
	var params = mux.Vars(r)

	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		utils.GetError(err, w)
		return
	}

	// prepare filter.
	filter := bson.M{"_id": id}

	deleteResult, err := database.Coll_product.DeleteOne(context.TODO(), filter)

	if err != nil {
		utils.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}
