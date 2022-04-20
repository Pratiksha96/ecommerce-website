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

func (pm *reviewManager) CreateReview(review models.Review, product models.Product, filterProduct primitive.M) (map[string]interface{}, error) {

	if review.Rating < 0 || review.Rating > 5 {
		return nil, errors.New("Product Rating cannot be negative or greater than 5")
	}
	product.Reviews = append(product.Reviews, &review)
	product.NumOfReviews = len(product.Reviews)
	var oldProduct models.Product
	avgRating := 0
	for _, reviewInstance := range product.Reviews {

		avgRating += reviewInstance.Rating
	}
	avgRating = avgRating / product.NumOfReviews

	product.Ratings = avgRating

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
			{"numOfReviews", product.NumOfReviews},
		}},
	}

	err := database.Coll_product.FindOneAndUpdate(context.TODO(), filterProduct, update).Decode(&oldProduct)

	if err != nil {
		return nil, err
	}

	ratingResponse := map[string]interface{}{"success": true, "message": "Review has been created"}
	return ratingResponse, nil
}

func (pm *reviewManager) GetProductReviews(id primitive.ObjectID) ([]*models.Review, error) {
	product := &models.Product{}
	filter := bson.M{"_id": id}
	err := database.Coll_product.FindOne(context.TODO(), filter).Decode(product)
	if err != nil {
		return nil, err
	}
	result := product.Reviews

	payload := result
	return payload, nil
}

func (pm *reviewManager) UpdateReview(review models.Review, product models.Product, filterProduct primitive.M) (map[string]interface{}, error) {
	if review.Rating < 0 || review.Rating > 5 {
		return nil, errors.New("Product Rating cannot be negative or greater than 5")
	}
	var oldProduct models.Product
	oldRating := 0
	for x := 0; x < len(product.Reviews); x++ {
		if product.Reviews[x].User.Email == review.User.Email {
			oldRating = product.Reviews[x].Rating
			product.Reviews[x] = &review
			break
		}
	}
	avgRating := product.Ratings
	product.Ratings = ((avgRating * product.NumOfReviews) - oldRating + review.Rating) / product.NumOfReviews

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
			{"numOfReviews", product.NumOfReviews},
		}},
	}

	err := database.Coll_product.FindOneAndUpdate(context.TODO(), filterProduct, update).Decode(&oldProduct)

	if err != nil {
		return nil, err
	}

	ratingResponse := map[string]interface{}{"success": true, "message": "Review has been updated"}
	return ratingResponse, nil
}

func (pm *reviewManager) DeleteReview(id primitive.ObjectID, email string) (map[string]interface{}, error) {

	product := &models.Product{}
	filter := bson.M{"_id": id}
	err := database.Coll_product.FindOne(context.TODO(), filter).Decode(product)
	if err != nil {
		return nil, err
	}

	indexToDelete := 0
	oldRating := 0
	for x := 0; x < len(product.Reviews); x++ {
		if product.Reviews[x].User.Email == email {
			oldRating = product.Reviews[x].Rating
			indexToDelete = x
			break
		}
	}
	avgRating := product.Ratings
	avgRating = ((avgRating * product.NumOfReviews) - oldRating) / (product.NumOfReviews - 1)

	newLength := 0
	for index := range product.Reviews {
		if indexToDelete != index {

			product.Reviews[newLength] = product.Reviews[index]
			newLength++
		}
	}

	// reslice the array to remove extra index
	newReview := product.Reviews[:newLength]
	newNumOfReviews := product.NumOfReviews - 1

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
