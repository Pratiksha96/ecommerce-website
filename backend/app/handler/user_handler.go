package handler

import (
	models "ecommerce-website/app/Models"
	"ecommerce-website/app/middleware"
	"ecommerce-website/app/utils"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	if errors := utils.UserRegisterValidation(user); len(errors) > 0 {
		log.Println("Received invalid json request!")
		err := map[string]interface{}{"success": false, "message": errors}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	} else {
		middleware.RegisterUser(user, w)
	}

}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	if errors := utils.UserLoginValidation(user); len(errors) > 0 {
		log.Println("Received invalid json request!")
		err := map[string]interface{}{"success": false, "message": errors}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	} else {
		middleware.LoginUser(user, w)
	}

}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	cookieToken, err := r.Cookie("token")

	if err != nil || len(cookieToken.Value) == 0 {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			utils.GetError(errors.New("received ErrNoCookie, you are already logged out"), w)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		utils.GetError(errors.New("please login to access this resource"), w)
		return
	}

	tokenStr := cookieToken.Value
	middleware.LogoutUser(tokenStr, w)
}
