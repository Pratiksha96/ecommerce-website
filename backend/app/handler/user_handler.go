package handler

import (
	"encoding/json"
	"net/http"

	models "ecommerce-website/app/Models"
	"ecommerce-website/app/middleware"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	middleware.RegisterUser(user, w)

}
