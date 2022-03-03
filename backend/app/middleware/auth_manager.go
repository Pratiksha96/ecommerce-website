package middleware

import (
	"context"
	models "ecommerce-website/app/models"
	"ecommerce-website/app/utils"
	"ecommerce-website/internal/database"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
)

func AuthenticateUser(w http.ResponseWriter, r *http.Request) (bool, string) {

	cookieToken, err := r.Cookie("token")

	if err != nil || len(cookieToken.Value) == 0 {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			utils.GetError(errors.New("Received ErrNoCookie, please login to access this resource"), w)
			return false, ""
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		utils.GetError(errors.New("Please login to access this resource"), w)
		return false, ""
	}

	tokenStr := cookieToken.Value

	claims := jwt.MapClaims{}

	//TODO to change jkey
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("ThisIsMySecretKey"), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			utils.GetError(errors.New("signature invalid"), w)
			return false, ""
		}
		w.WriteHeader(http.StatusBadRequest)
		utils.GetError(errors.New("received error while validating token"), w)
		return false, ""
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		utils.GetError(errors.New("token is invalid"), w)
		return false, ""
	}

	tknClaims, _ := tkn.Claims.(jwt.MapClaims)
	email, _ := tknClaims["email"].(string)
	println("User email received is: ", email)
	return true, email
}

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
