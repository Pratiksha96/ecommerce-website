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

	fmt.Println("User registered: ", registeredUser.InsertedID)
	user.Password = ""
	json.NewEncoder(w).Encode(user)

}
