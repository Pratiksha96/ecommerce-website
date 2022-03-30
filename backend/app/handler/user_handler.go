package handler

import (
	"ecommerce-website/app/manager"
	models "ecommerce-website/app/models"
	"ecommerce-website/app/utils"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
			manager.RegisterUser(user, w)
		}
	}
}

func LoginUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
			manager.LoginUser(user, w)
		}
	}
}

func LogoutUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		manager.LogoutUser(tokenStr, w)
	}
}

func UserDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		ctx := r.Context()
		email := ctx.Value("email").(string)
		manager.GetUserDetails(email, w)
	}
}

func UpdatePassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		var requestbody interface{}
		buffer, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Panic(err)
		}
		r.Body.Close()
		json.Unmarshal(buffer, &requestbody)
		body := requestbody.(map[string]interface{})

		ctx := r.Context()
		email := ctx.Value("email").(string)
		manager.UpdatePassword(email, body, w)
	}
}

func UpdateProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		var requestbody interface{}
		buffer, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Panic(err)
		}
		r.Body.Close()
		json.Unmarshal(buffer, &requestbody)
		body := requestbody.(map[string]interface{})

		ctx := r.Context()
		email := ctx.Value("email").(string)
		manager.UpdateProfile(email, body, w)
	}
}

func GetAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		ctx := r.Context()
		email := ctx.Value("email").(string)

		manager.GetAllUsers("admin", email, w)

	}
}

func GetSingleUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var params = mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil {
			utils.GetErrorWithStatus(errors.New("invalid object id"), w, http.StatusUnprocessableEntity)
			return
		}
		ctx := r.Context()
		email := ctx.Value("email").(string)

		manager.GetSingleUser("admin", email, id, w)

	}
}
