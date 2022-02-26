package middleware

import (
	"context"
	models "ecommerce-website/app/Models"
	"ecommerce-website/app/utils"
	"ecommerce-website/internal/database"

	"golang.org/x/crypto/bcrypt"

	"encoding/json"
	"fmt"
	"net/http"

	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
)

const SecretKey = "ThisIsMySecretKey"

func RegisterUser(user models.User, w http.ResponseWriter) {

	user.GetBSON()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)
	registeredUser, err := database.Coll_user.InsertOne(context.TODO(), user)

	if err != nil {
		utils.GetError(err, w)
		return
	}

	fmt.Println("User registered: ", registeredUser.InsertedID)
	user.Password = ""
	json.NewEncoder(w).Encode(user)

}

func LoginUser(user models.User, w http.ResponseWriter) {

	var storedUser models.User
	filter := bson.M{"email": user.Email}
	err := database.Coll_user.FindOne(context.TODO(), filter).Decode(&storedUser)

	if err != nil {
		utils.GetError(errors.New("no such user present"), w)
		return
	}

	passwordErr := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if passwordErr != nil {
		utils.GetError(errors.New("password mismatched"), w)
		return
	}

	// Declaring the expiration time of the user  token here
	// We have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute).Unix()

	// Create the JWT claims, which includes the user email and expiry time
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Email,
		ExpiresAt: expirationTime,
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		utils.GetError(err, w)
		return
	}

	fmt.Println("User logged in successfully: ", user.Email)
	json.NewEncoder(w).Encode(token)
}
