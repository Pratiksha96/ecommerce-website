package handler

import (
	"encoding/json"
	"log"
	"net/http"

	models "ecommerce-website/app/Models"
	"ecommerce-website/app/middleware"
	"ecommerce-website/app/utils"
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
