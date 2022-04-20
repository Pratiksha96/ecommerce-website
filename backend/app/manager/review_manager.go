package manager

import (
	"context"
	models "ecommerce-website/app/models"
	"ecommerce-website/internal/database"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReviewManager interface {
	CreateReview(review models.Review, product models.Product, filterProduct primitive.M) (map[string]interface{}, error)
	GetProductReviews(id primitive.ObjectID) ([]*models.Review, error)
	UpdateReview(review models.Review, product models.Product, filterProduct primitive.M) (map[string]interface{}, error)
	DeleteReview(id primitive.ObjectID, email string) (map[string]interface{}, error)
}

type reviewManager struct{}

func NewReviewManager() ReviewManager {
	return &reviewManager{}
}

//function that creates the review for a particular product
func (pm *reviewManager) CreateReview(review models.Review, product models.Product, filterProduct primitive.M) (map[string]interface{}, error) {

	//Checking for validiy of rating in review
	if review.Rating < 0 || review.Rating > 5 {
		return nil, errors.New("Product Rating cannot be negative or greater than 5")
	}
	//adding review in product
	product.Reviews = append(product.Reviews, &review)
	product.NumOfReviews = len(product.Reviews)
	//calculating average rating by getting total of previous rating, adding bew rating and dividing it by total updated reviews
	avgRating := product.Ratings
	product.Ratings = ((avgRating * product.NumOfReviews) + review.Rating) / product.NumOfReviews

	product.Ratings = avgRating

	//updating the review and other changed information in DB
	result, err := database.Coll_product.UpdateOne(
		context.TODO(),
		filterProduct,
		bson.D{
			{"$set", bson.D{{"reviews", product.Reviews},
				{"numOfReviews", product.NumOfReviews},
				{"ratings", product.Ratings}}},
		},
	)

	if err != nil {
		return nil, err
	}
	log.Println("Following number of users updated ", result.ModifiedCount)
	ratingResponse := map[string]interface{}{"success": true, "message": "Review has been created"}
	return ratingResponse, nil
}

//function that gets all reviews for a particular product
func (pm *reviewManager) GetProductReviews(id primitive.ObjectID) ([]*models.Review, error) {
	product := &models.Product{}
	filter := bson.M{"_id": id}
	err := database.Coll_product.FindOne(context.TODO(), filter).Decode(product)
	if err != nil {
		return nil, err
	}
	//getting reviews of product
	result := product.Reviews
	//setting the result in payload
	payload := result
	return payload, nil
}

//function that updates a review made by user for a particular product
func (pm *reviewManager) UpdateReview(review models.Review, product models.Product, filterProduct primitive.M) (map[string]interface{}, error) {

	//Checking for validiy of rating in review
	if review.Rating < 0 || review.Rating > 5 {
		return nil, errors.New("Product Rating cannot be negative or greater than 5")
	}

	oldRating := 0
	//Traversing through all the reviews and getting the review that this particular user made
	for counter := 0; counter < len(product.Reviews); counter++ {
		if product.Reviews[counter].User.Email == review.User.Email {
			//saving the old rating made by user, for updating average review
			oldRating = product.Reviews[counter].Rating
			product.Reviews[counter] = &review
			break
		}
	}
	//Calcuating avg rating by subtracting old rating and adding new updated rating
	avgRating := product.Ratings
	product.Ratings = ((avgRating * product.NumOfReviews) - oldRating + review.Rating) / product.NumOfReviews
	//updating the new review
	result, err := database.Coll_product.UpdateOne(
		context.TODO(),
		filterProduct,
		bson.D{
			{"$set", bson.D{{"reviews", product.Reviews},
				{"ratings", product.Ratings}}},
		},
	)
	if err != nil {
		return nil, err
	}
	log.Println("Following number of users updated ", result.ModifiedCount)
	ratingResponse := map[string]interface{}{"success": true, "message": "Review has been updated"}
	return ratingResponse, nil
}

//function that deletes the review made by a user for a product
func (pm *reviewManager) DeleteReview(id primitive.ObjectID, email string) (map[string]interface{}, error) {

	product := &models.Product{}
	filter := bson.M{"_id": id}
	err := database.Coll_product.FindOne(context.TODO(), filter).Decode(product)
	if err != nil {
		return nil, err
	}

	indexToDelete := 0
	oldRating := 0
	//finding index of review that was created by this particular user
	for counter := 0; counter < len(product.Reviews); counter++ {
		if product.Reviews[counter].User.Email == email {
			oldRating = product.Reviews[counter].Rating
			indexToDelete = counter
			break
		}
	}
	//calcualting the new avg rating by subtracting the deleted reviews rating from total
	avgRating := product.Ratings
	avgRating = ((avgRating * product.NumOfReviews) - oldRating) / (product.NumOfReviews - 1)

	newLength := 0
	//Copy the elements before the indexTodelete
	for index := range product.Reviews {
		if indexToDelete != index {

			product.Reviews[newLength] = product.Reviews[index]
			newLength++
		}
	}

	// reslice the array to remove extra index, that is shift each element by one index
	newReview := product.Reviews[:newLength]
	newNumOfReviews := product.NumOfReviews - 1
	// update the new rating and deleted review in DB
	result, err := database.Coll_product.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"reviews", newReview},
				{"numOfReviews", newNumOfReviews},
				{"ratings", avgRating}}},
		},
	)
	if err != nil {
		return nil, errors.New("Error while deleting product review")
	}
	log.Println("Following number of users updated ", result.ModifiedCount)

	deleteResponse := map[string]interface{}{"success": true, "message": "Review has been successfully deleted"}
	return deleteResponse, nil
}
