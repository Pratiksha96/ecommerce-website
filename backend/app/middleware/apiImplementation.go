package middleware

import (
	"context"
	models "ecommerce-website/app/Models"
	"ecommerce-website/app/utils"
	"ecommerce-website/internal/database"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
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
func GetProduct(id primitive.ObjectID, w http.ResponseWriter) {

	var product models.Product
	filter := bson.M{"_id": id}
	err := database.Coll_product.FindOne(context.TODO(), filter).Decode(&product)

	if err != nil {
		utils.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(product)

}

func GetAllProducts(w http.ResponseWriter) {
	cur, err := database.Coll_product.Find(context.Background(), bson.D{{}})
	if err != nil {
		utils.GetError(err, w)
		return
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			utils.GetError(err, w)
			return
		}
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		utils.GetError(err, w)
		return
	} else if len(results) == 0 {
		utils.GetError(errors.New("product list is empty"), w)
		return
	}

	cur.Close(context.Background())
	payload := results
	json.NewEncoder(w).Encode(payload)

}

func SearchProducts(w http.ResponseWriter, query url.Values) {

	keyword, keywordPresent := query["keyword"]
	categoryType, categoryPresent := query["category"]
	priceMinRange, priceMinPresent := query["priceMin"]
	priceMaxRange, priceMaxPresent := query["priceMax"]

	if (keywordPresent || len(keyword) > 0) && (categoryPresent || len(categoryType) > 0) {
		if (priceMinPresent || len(priceMinRange) > 0) && (priceMaxPresent || len(priceMaxRange) > 0) {
			name := keyword[0]
			category := categoryType[0]
			priceMin := priceMinRange[0]
			priceMinNum, _ := strconv.Atoi(priceMin)

			priceMax := priceMaxRange[0]
			priceMaxNum, _ := strconv.Atoi(priceMax)

			filter := bson.D{{"name", primitive.Regex{Pattern: name, Options: ""}}, {"category", category}, {"price", bson.D{{"$gte", priceMinNum}, {"$lte", priceMaxNum}}}}
			GetData(filter, w)
		} else if priceMaxPresent || len(priceMaxRange) > 0 {
			name := keyword[0]
			category := categoryType[0]
			priceMax := priceMaxRange[0]
			priceMaxNum, _ := strconv.Atoi(priceMax)

			filter := bson.D{{"name", primitive.Regex{Pattern: name, Options: ""}}, {"category", category}, {"price", bson.D{{"$gte", 1}, {"$lte", priceMaxNum}}}}
			GetData(filter, w)
		} else if priceMinPresent || len(priceMinRange) > 0 {
			name := keyword[0]
			category := categoryType[0]
			priceMin := priceMinRange[0]
			priceMinNum, _ := strconv.Atoi(priceMin)

			filter := bson.D{{"name", primitive.Regex{Pattern: name, Options: ""}}, {"category", category}, {"price", bson.D{{"$gte", priceMinNum}, {"$lte", 99999999}}}}
			GetData(filter, w)
		} else {
			name := keyword[0]
			category := categoryType[0]
			filter := bson.D{{"name", primitive.Regex{Pattern: name, Options: ""}}, {"category", category}}
			GetData(filter, w)
		}

	} else if keywordPresent || len(keyword) > 0 {
		if (priceMinPresent || len(priceMinRange) > 0) && (priceMaxPresent || len(priceMaxRange) > 0) {
			name := keyword[0]
			priceMin := priceMinRange[0]
			priceMinNum, _ := strconv.Atoi(priceMin)

			priceMax := priceMaxRange[0]
			priceMaxNum, _ := strconv.Atoi(priceMax)

			filter := bson.D{{"name", primitive.Regex{Pattern: name, Options: ""}}, {"price", bson.D{{"$gte", priceMinNum}, {"$lte", priceMaxNum}}}}
			GetData(filter, w)
		} else if priceMaxPresent || len(priceMaxRange) > 0 {
			name := keyword[0]
			priceMax := priceMaxRange[0]
			priceMaxNum, _ := strconv.Atoi(priceMax)

			filter := bson.D{{"name", primitive.Regex{Pattern: name, Options: ""}}, {"price", bson.D{{"$gte", 1}, {"$lte", priceMaxNum}}}}
			GetData(filter, w)
		} else if priceMinPresent || len(priceMinRange) > 0 {
			name := keyword[0]
			priceMin := priceMinRange[0]
			priceMinNum, _ := strconv.Atoi(priceMin)

			filter := bson.D{{"name", primitive.Regex{Pattern: name, Options: ""}}, {"price", bson.D{{"$gte", priceMinNum}, {"$lte", 99999999}}}}
			GetData(filter, w)
		} else {
			name := keyword[0]
			filter := bson.D{{"name", primitive.Regex{Pattern: name, Options: ""}}}
			GetData(filter, w)
		}
	} else if categoryPresent || len(categoryType) > 0 {
		category := categoryType[0]
		filter := bson.D{{"category", category}}
		GetData(filter, w)
	} else if (priceMinPresent || len(priceMinRange) > 0) && (priceMaxPresent || len(priceMaxRange) > 0) {
		priceMin := priceMinRange[0]
		priceMinNum, _ := strconv.Atoi(priceMin)

		priceMax := priceMaxRange[0]
		priceMaxNum, _ := strconv.Atoi(priceMax)

		filter := bson.D{{"price", bson.D{{"$gte", priceMinNum}, {"$lte", priceMaxNum}}}}
		GetData(filter, w)
	} else if priceMaxPresent || len(priceMaxRange) > 0 {
		priceMax := priceMaxRange[0]
		priceMaxNum, _ := strconv.Atoi(priceMax)

		filter := bson.D{{"price", bson.D{{"$gte", 1}, {"$lte", priceMaxNum}}}}
		GetData(filter, w)
	} else if priceMinPresent || len(priceMinRange) > 0 {
		priceMin := priceMinRange[0]
		priceMinNum, _ := strconv.Atoi(priceMin)

		filter := bson.D{{"price", bson.D{{"$gte", priceMinNum}, {"$lte", 99999999}}}}
		GetData(filter, w)
	} else {
		filter := bson.D{{}}
		GetData(filter, w)
	}
}

func GetData(filter bson.D, w http.ResponseWriter) {
	cur, err := database.Coll_product.Find(context.Background(), filter)

	if err != nil {
		utils.GetError(err, w)
		return
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			utils.GetError(err, w)
			return
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		utils.GetError(err, w)
		return
	} else if len(results) == 0 {
		utils.GetError(errors.New("product list is empty"), w)
		return
	}

	cur.Close(context.Background())
	payload := results
	json.NewEncoder(w).Encode(payload)
}

func UpdateProduct(id primitive.ObjectID, product models.Product, w http.ResponseWriter) {

	filter := bson.M{"_id": id}

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

func DeleteProduct(id primitive.ObjectID, w http.ResponseWriter) {

	filter := bson.M{"_id": id}

	deleteResult, err := database.Coll_product.DeleteOne(context.TODO(), filter)

	if err != nil {
		utils.GetError(err, w)
		return
	} else if deleteResult.DeletedCount == 0 {
		utils.GetError(errors.New("no such document present"), w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)

}
