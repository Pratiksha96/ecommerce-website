package middleware

import (
	//"context"
	"context"
	models "ecommerce-website/app/Models"
	"ecommerce-website/app/utils"
	"ecommerce-website/internal/database"

	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(product models.Product, w http.ResponseWriter) {

	insertedProduct, err := database.Coll_product.InsertOne(context.TODO(), product)

	if err != nil {
		utils.GetError(err, w)
		return
	}

	fmt.Println("Product Inserted", insertedProduct.InsertedID)

	json.NewEncoder(w).Encode(product)

}
