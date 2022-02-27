package middleware

import (
	"context"
	models "ecommerce-website/app/Models"
	"ecommerce-website/app/utils"
	"ecommerce-website/internal/database"
	"encoding/json"

	"golang.org/x/crypto/bcrypt"

	"fmt"
	"net/http"

	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

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

	token, err := user.GetJwtToken()
	if err != nil {
		utils.GetError(err, w)
		return
	}
	fmt.Println("User registered successfully with id: ", user.Email, registeredUser.InsertedID)
	utils.StoreUserToken(token, w)
	tokenResponse := map[string]interface{}{"success": true, "token": token}
	json.NewEncoder(w).Encode(tokenResponse)
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

	token, err := user.GetJwtToken()
	if err != nil {
		utils.GetError(err, w)
		return
	}

	fmt.Println("User logged in successfully: ", user.Email)
	utils.StoreUserToken(token, w)

	tokenResponse := map[string]interface{}{"success": true, "token": token}
	json.NewEncoder(w).Encode(tokenResponse)
}
