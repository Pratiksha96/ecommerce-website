package handler

import (
	"context"
	"encoding/json"
	"errors"
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

//handler that will handle create review API request
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
		// fetching review data from request body
		body := requestbody.(map[string]interface{})

		ctx := r.Context()
		email := ctx.Value("email").(string)

		var storedUser models.User
		filter := bson.M{"email": email}
		//getting user details from DB who is creating review
		err = database.Coll_user.FindOne(context.TODO(), filter).Decode(&storedUser)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		ratingString := body["rating"].(string)
		ratingInt, err := strconv.Atoi(ratingString)
		proudctIdString := body["productId"].(string)
		//getting product id from the URL parameters
		productId, err := primitive.ObjectIDFromHex(proudctIdString)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		//creating filter to get product from DB
		filterProduct := bson.M{"_id": productId}
		err = database.Coll_product.FindOne(context.TODO(), filterProduct).Decode(&product)
		if err != nil {
			utils.GetError(err, w)
			return
		}

		var review models.Review
		// creating Review object that will be added in Product
		review.User = storedUser
		review.Name = storedUser.Name
		review.Rating = ratingInt
		review.Comment = body["comment"].(string)
		//calling review manager to add review in product
		response, err := reviewManager.CreateReview(review, product, filterProduct)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

//handler that will handle get reviews for a particular product API request
func GetProductReviews(reviewManager manager.ReviewManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var params = mux.Vars(r)
		//getting product id from the URL parameters
		id, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil {
			utils.GetErrorWithStatus(errors.New("invalid product id"), w, http.StatusUnprocessableEntity)
			return
		}
		//calling review manager to get all reviews for a particular product
		payload, err := reviewManager.GetProductReviews(id)
		if err != nil {
			utils.GetErrorWithStatus(errors.New("invalid product id"), w, http.StatusUnprocessableEntity)
			return
		}
		json.NewEncoder(w).Encode(payload)
	}
}

//handler that will handle update review API request
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
		// fetching review data from request body
		body := requestbody.(map[string]interface{})

		ctx := r.Context()
		email := ctx.Value("email").(string)

		var storedUser models.User
		filter := bson.M{"email": email}
		//getting user details from DB who is updating review
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
		// getting product in which review is to be updated
		err = database.Coll_product.FindOne(context.TODO(), filterProduct).Decode(&product)
		if err != nil {
			utils.GetError(err, w)
			return
		}

		var review models.Review
		// creating Review object that is to be upated in Product
		review.User = storedUser
		review.Name = storedUser.Name
		review.Rating = ratingInt
		review.Comment = body["comment"].(string)
		//calling review manager to update the review formed aboce
		response, err := reviewManager.UpdateReview(review, product, filterProduct)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

//handler that will handle delete review API request
func DeleteReview(reviewManager manager.ReviewManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var params = mux.Vars(r)
		//getting product id from URL parameters
		id, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil || params["id"] == "" {
			utils.GetErrorWithStatus(errors.New("invalid object id"), w, http.StatusUnprocessableEntity)
			return
		}
		ctx := r.Context()
		email := ctx.Value("email").(string)
		//calling reviewManager to delete the review that this user created
		deleteResponse, err := reviewManager.DeleteReview(id, email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(deleteResponse)
	}
}
