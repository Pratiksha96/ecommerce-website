package middleware

import (
	"context"
	models "ecommerce-website/app/Models"
	"ecommerce-website/app/utils"
	"ecommerce-website/internal/database"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func AuthorizeUser(w http.ResponseWriter, role string, email string) bool {

	var user models.User
	userFilter := bson.M{"email": email}
	userErr := database.Coll_user.FindOne(context.TODO(), userFilter).Decode(&user)

	if userErr != nil {
		utils.GetError(userErr, w)
		return false
	}

	if role == "admin" && (role != user.Role) {
		utils.GetError(errors.New("sorry, you don't have access to this resource"), w)
		return false
	}
	return true
}
