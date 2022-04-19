package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"ecommerce-website/app/manager"
	models "ecommerce-website/app/models"
	"ecommerce-website/app/utils"
	"ecommerce-website/internal/database"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateReview(reviewManager manager.ReviewManager) http.HandlerFunc {
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
			utils.GetError(err, w)
			return
		}
		ratingString := body["rating"].(string)
		ratingInt, err := strconv.Atoi(ratingString)
		proudctIdString := body["productId"].(string)
		productId, err := primitive.ObjectIDFromHex(proudctIdString)
		if err != nil {
			utils.GetError(err, w)
			return
		}

		filterProduct := bson.M{"_id": productId}
		err = database.Coll_product.FindOne(context.TODO(), filterProduct).Decode(&product)
		if err != nil {
			utils.GetError(err, w)
			return
		}

		var review models.Review

		review.User = storedUser
		review.Name = storedUser.Name
		review.Rating = ratingInt
		review.Comment = body["comment"].(string)

		response, err := reviewManager.CreateReview(review, product, filterProduct)
		if err != nil {
			fmt.Print("Caught error on line 214")
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

func GetProductReviews(reviewManager manager.ReviewManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var params = mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil {
			utils.GetErrorWithStatus(errors.New("invalid product id"), w, http.StatusUnprocessableEntity)
			return
		}
		payload, err := reviewManager.GetProductReviews(id)
		if err != nil {
			utils.GetErrorWithStatus(errors.New("invalid product id"), w, http.StatusUnprocessableEntity)
			return
		}
		json.NewEncoder(w).Encode(payload)
	}
}

func UpdateReview(reviewManager manager.ReviewManager) http.HandlerFunc {
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

		response, err := reviewManager.UpdateReview(review, product, filterProduct)
		if err != nil {
			fmt.Print("Caught error on line 214")
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

func DeleteReview(reviewManager manager.ReviewManager) http.HandlerFunc {
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
		deleteResponse, err := reviewManager.DeleteReview(id, email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(deleteResponse)
	}
}
