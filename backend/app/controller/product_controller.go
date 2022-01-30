package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	models "ecommerce-website/app/Models"

	"ecommerce-website/internal/database"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	fmt.Println("this is produc + ", r.Body)
	insertResult, err := database.Coll_product.InsertOne(context.Background(), product)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)

	// fmt.Println(task, r.Body)

	json.NewEncoder(w).Encode(product)
}
