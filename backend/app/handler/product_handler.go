package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"ecommerce-website/app/manager"
	models "ecommerce-website/app/models"
	"ecommerce-website/app/utils"

	"ecommerce-website/internal/database"
	"log"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
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

		if err := utils.Validate(product); len(err) > 0 {
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

		if err := utils.Validate(product); len(err) > 0 {
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

func CreateReview(productManager manager.ProductManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		var product models.Product

		var requestbody interface{}
		buffer, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Panic(err)
		}
		r.Body.Close()
		json.Unmarshal(buffer, &requestbody)
		body := requestbody.(map[string]interface{})

		ctx := r.Context()
		email := ctx.Value("email").(string)

		var storedUser models.User
		filter := bson.M{"email": email}
		err = database.Coll_user.FindOne(context.TODO(), filter).Decode(&storedUser)
		if err != nil {
			fmt.Print("Caught error on line 175")
			utils.GetError(err, w)
			return
		}
		ratingString := body["rating"].(string)
		ratingInt, err := strconv.Atoi(ratingString)
		proudctIdString := body["productId"].(string)
		productId, err := primitive.ObjectIDFromHex(proudctIdString)
		if err != nil {
			fmt.Print("Caught error on line 184")
			utils.GetError(err, w)
			return
		}

		filterProduct := bson.M{"_id": productId}
		err = database.Coll_product.FindOne(context.TODO(), filterProduct).Decode(&product)
		if err != nil {
			fmt.Print("Caught error on line 192")
			utils.GetError(err, w)
			return
		}

		var review models.Review

		review.User = storedUser
		review.Name = storedUser.Name
		review.Rating = ratingInt
		review.Comment = body["comment"].(string)

		// if errors := utils.Validate(product); len(errors) > 0 {
		// 	log.Println("Received invalid json request!")
		// 	err := map[string]interface{}{"success": false, "message": errors}
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	json.NewEncoder(w).Encode(err)
		// 	return
		// }

		response, err := productManager.CreateReview(review, product, filterProduct)
		if err != nil {
			fmt.Print("Caught error on line 214")
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

func GetProductReviews(productManager manager.ProductManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var params = mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil {
			utils.GetErrorWithStatus(errors.New("invalid object id"), w, http.StatusUnprocessableEntity)
			return
		}
		payload, err := productManager.GetProductReviews(id)
		if err != nil {
			utils.GetErrorWithStatus(errors.New("invalid object id"), w, http.StatusUnprocessableEntity)
			return
		}
		json.NewEncoder(w).Encode(payload)
	}
}

func UpdateReview(productManager manager.ProductManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		var product models.Product

		var requestbody interface{}
		buffer, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Panic(err)
		}
		r.Body.Close()
		json.Unmarshal(buffer, &requestbody)
		body := requestbody.(map[string]interface{})

		ctx := r.Context()
		email := ctx.Value("email").(string)

		var storedUser models.User
		filter := bson.M{"email": email}
		err = database.Coll_user.FindOne(context.TODO(), filter).Decode(&storedUser)
		if err != nil {
			fmt.Print("Caught error on line 175")
			utils.GetError(err, w)
			return
		}
		ratingString := body["rating"].(string)
		ratingInt, err := strconv.Atoi(ratingString)
		proudctIdString := body["productId"].(string)
		productId, err := primitive.ObjectIDFromHex(proudctIdString)
		if err != nil {
			fmt.Print("Caught error on line 184")
			utils.GetError(err, w)
			return
		}

		filterProduct := bson.M{"_id": productId}
		err = database.Coll_product.FindOne(context.TODO(), filterProduct).Decode(&product)
		if err != nil {
			fmt.Print("Caught error on line 192")
			utils.GetError(err, w)
			return
		}

		var review models.Review

		review.User = storedUser
		review.Name = storedUser.Name
		review.Rating = ratingInt
		review.Comment = body["comment"].(string)

		response, err := productManager.UpdateReview(review, product, filterProduct)
		if err != nil {
			fmt.Print("Caught error on line 214")
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}
