package manager

import (
	"context"
	models "ecommerce-website/app/models"
	"ecommerce-website/app/utils"
	"ecommerce-website/internal/database"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type ProductManager interface {
	GetProduct(id primitive.ObjectID, email string) (*models.Product, error)
	CreateProduct(product models.Product, role string, email string) (*models.Product, error)
	GetAllProducts(email string) ([]primitive.M, error)
	UpdateProduct(id primitive.ObjectID, product models.Product, role string, email string) (*models.Product, error)
	DeleteProduct(id primitive.ObjectID, role string, email string) (map[string]interface{}, error)
}

type productManager struct{}

func NewProductManager() ProductManager {
	return &productManager{}
}

func (pm *productManager) CreateProduct(product models.Product, role string, email string) (*models.Product, error) {
	err := authorizeUser(role, email)
	if err != nil {
		return nil, err
	}

	insertedProduct, err := database.Coll_product.InsertOne(context.TODO(), product)
	if err != nil {
		return nil, err
	}

	log.Println("Product Inserted", insertedProduct.InsertedID)
	return &product, nil
}

func (pm *productManager) GetProduct(id primitive.ObjectID, email string) (*models.Product, error) {
	product := &models.Product{}
	filter := bson.M{"_id": id}
	err := database.Coll_product.FindOne(context.TODO(), filter).Decode(product)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (pm *productManager) GetAllProducts(email string) ([]primitive.M, error) {
	cur, err := database.Coll_product.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	} else if len(results) == 0 {
		return nil, errors.New("Product list is empty")
	}

	cur.Close(context.Background())
	payload := results
	return payload, nil
}

func SearchProducts(w http.ResponseWriter, query url.Values, email string) {

	keyword, keywordPresent := query["keyword"]
	categoryType, categoryPresent := query["category"]
	priceMinRange, priceMinPresent := query["priceMin"]
	priceMaxRange, priceMaxPresent := query["priceMax"]

	currentPage, isPagePresent := query["page"]
	var resultsPerPage int64 = 10
	var skips int64 = 0

	if isPagePresent || len(currentPage) > 0 {
		current := currentPage[0]
		currentPageNum, _ := strconv.Atoi(current)
		skips = (resultsPerPage) * int64((currentPageNum - 1))
	}

	if (keywordPresent || len(keyword) > 0) && (categoryPresent || len(categoryType) > 0) {
		if (priceMinPresent || len(priceMinRange) > 0) && (priceMaxPresent || len(priceMaxRange) > 0) {
			name := keyword[0]
			category := categoryType[0]
			priceMin := priceMinRange[0]
			priceMinNum, _ := strconv.Atoi(priceMin)

			priceMax := priceMaxRange[0]
			priceMaxNum, _ := strconv.Atoi(priceMax)

			filter := bson.D{{"name", primitive.Regex{Pattern: name, Options: ""}}, {"category", category}, {"price", bson.D{{"$gte", priceMinNum}, {"$lte", priceMaxNum}}}}
			GetFilteredProducts(filter, w, resultsPerPage, skips)
		} else if priceMaxPresent || len(priceMaxRange) > 0 {
			name := keyword[0]
			category := categoryType[0]
			priceMax := priceMaxRange[0]
			priceMaxNum, _ := strconv.Atoi(priceMax)

			filter := bson.D{{"name", primitive.Regex{Pattern: name, Options: ""}}, {"category", category}, {"price", bson.D{{"$gte", 1}, {"$lte", priceMaxNum}}}}
			GetFilteredProducts(filter, w, resultsPerPage, skips)
		} else if priceMinPresent || len(priceMinRange) > 0 {
			name := keyword[0]
			category := categoryType[0]
			priceMin := priceMinRange[0]
			priceMinNum, _ := strconv.Atoi(priceMin)

			filter := bson.D{{"name", primitive.Regex{Pattern: name, Options: ""}}, {"category", category}, {"price", bson.D{{"$gte", priceMinNum}, {"$lte", 99999999}}}}
			GetFilteredProducts(filter, w, resultsPerPage, skips)
		} else {
			name := keyword[0]
			category := categoryType[0]
			filter := bson.D{{"name", primitive.Regex{Pattern: name, Options: ""}}, {"category", category}}
			GetFilteredProducts(filter, w, resultsPerPage, skips)
		}

	} else if keywordPresent || len(keyword) > 0 {
		if (priceMinPresent || len(priceMinRange) > 0) && (priceMaxPresent || len(priceMaxRange) > 0) {
			name := keyword[0]
			priceMin := priceMinRange[0]
			priceMinNum, _ := strconv.Atoi(priceMin)

			priceMax := priceMaxRange[0]
			priceMaxNum, _ := strconv.Atoi(priceMax)

			filter := bson.D{{"name", primitive.Regex{Pattern: name, Options: ""}}, {"price", bson.D{{"$gte", priceMinNum}, {"$lte", priceMaxNum}}}}
			GetFilteredProducts(filter, w, resultsPerPage, skips)
		} else if priceMaxPresent || len(priceMaxRange) > 0 {
			name := keyword[0]
			priceMax := priceMaxRange[0]
			priceMaxNum, _ := strconv.Atoi(priceMax)

			filter := bson.D{{"name", primitive.Regex{Pattern: name, Options: ""}}, {"price", bson.D{{"$gte", 1}, {"$lte", priceMaxNum}}}}
			GetFilteredProducts(filter, w, resultsPerPage, skips)
		} else if priceMinPresent || len(priceMinRange) > 0 {
			name := keyword[0]
			priceMin := priceMinRange[0]
			priceMinNum, _ := strconv.Atoi(priceMin)

			filter := bson.D{{"name", primitive.Regex{Pattern: name, Options: ""}}, {"price", bson.D{{"$gte", priceMinNum}, {"$lte", 99999999}}}}
			GetFilteredProducts(filter, w, resultsPerPage, skips)
		} else {
			name := keyword[0]
			filter := bson.D{{"name", primitive.Regex{Pattern: name, Options: ""}}}
			GetFilteredProducts(filter, w, resultsPerPage, skips)
		}
	} else if categoryPresent || len(categoryType) > 0 {
		category := categoryType[0]
		filter := bson.D{{"category", category}}
		GetFilteredProducts(filter, w, resultsPerPage, skips)
	} else if (priceMinPresent || len(priceMinRange) > 0) && (priceMaxPresent || len(priceMaxRange) > 0) {
		priceMin := priceMinRange[0]
		priceMinNum, _ := strconv.Atoi(priceMin)

		priceMax := priceMaxRange[0]
		priceMaxNum, _ := strconv.Atoi(priceMax)

		filter := bson.D{{"price", bson.D{{"$gte", priceMinNum}, {"$lte", priceMaxNum}}}}
		GetFilteredProducts(filter, w, resultsPerPage, skips)
	} else if priceMaxPresent || len(priceMaxRange) > 0 {
		priceMax := priceMaxRange[0]
		priceMaxNum, _ := strconv.Atoi(priceMax)

		filter := bson.D{{"price", bson.D{{"$gte", 1}, {"$lte", priceMaxNum}}}}
		GetFilteredProducts(filter, w, resultsPerPage, skips)
	} else if priceMinPresent || len(priceMinRange) > 0 {
		priceMin := priceMinRange[0]
		priceMinNum, _ := strconv.Atoi(priceMin)

		filter := bson.D{{"price", bson.D{{"$gte", priceMinNum}, {"$lte", 99999999}}}}
		GetFilteredProducts(filter, w, resultsPerPage, skips)
	} else {
		filter := bson.D{{}}
		GetFilteredProducts(filter, w, resultsPerPage, skips)
	}
}

func GetFilteredProducts(filter bson.D, w http.ResponseWriter, resultsPerPage int64, skips int64) {

	opts := options.FindOptions{
		Skip:  &skips,
		Limit: &resultsPerPage,
	}
	cur, err := database.Coll_product.Find(context.TODO(), filter, &opts)

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

func (pm *productManager) UpdateProduct(id primitive.ObjectID, product models.Product, role string, email string) (*models.Product, error) {

	err := authorizeUser(role, email)
	if err != nil {
		return nil, err
	}
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

	err = database.Coll_product.FindOneAndUpdate(context.TODO(), filter, update).Decode(&oldProduct)

	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (pm *productManager) DeleteProduct(id primitive.ObjectID, role string, email string) (map[string]interface{}, error) {
	err := authorizeUser(role, email)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": id}
	deleteResult, err := database.Coll_product.DeleteOne(context.TODO(), filter)

	if err != nil {
		return nil, err
	} else if deleteResult.DeletedCount == 0 {
		return nil, errors.New("no such document present")
	}

	deleteResponse := map[string]interface{}{"success": true, "message": "document has been successfully deleted"}
	return deleteResponse, nil
}

func authorizeUser(role string, email string) error {
	var user models.User
	userFilter := bson.M{"email": email}
	userErr := database.Coll_user.FindOne(context.TODO(), userFilter).Decode(&user)

	if userErr != nil {
		return userErr
	}

	if role == "admin" && (role != user.Role) {
		return errors.New("sorry, you don't have access to this resource")
	}
	return nil
}
