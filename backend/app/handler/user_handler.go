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

func RegisterUser(userManager manager.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		var user models.User
		_ = json.NewDecoder(r.Body).Decode(&user)

		if errs := utils.UserRegisterValidation(user); len(errs) > 0 {
			utils.GetErrorWithStatus(errors.New("Received invalid json request!"), w, http.StatusBadRequest)
			return
		}
		tokenResponse, err := userManager.RegisterUser(user)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		utils.StoreUserToken(tokenResponse.Token, w)
		json.NewEncoder(w).Encode(tokenResponse)
	}
}

func LoginUser(userManager manager.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		var user models.User
		_ = json.NewDecoder(r.Body).Decode(&user)

		if errs := utils.UserLoginValidation(user); len(errs) > 0 {
			utils.GetErrorWithStatus(errors.New("Received invalid json request!"), w, http.StatusBadRequest)
			return
		}
		tokenResponse, err := userManager.LoginUser(user)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		utils.StoreUserToken(tokenResponse.Token, w)
		json.NewEncoder(w).Encode(tokenResponse)
	}
}

func LogoutUser(userManager manager.UserManager) http.HandlerFunc {
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
				utils.GetError(errors.New("Received ErrNoCookie, you are already logged out"), w)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			utils.GetError(errors.New("Please login to access this resource"), w)
			return
		}

		tokenStr := cookieToken.Value
		utils.DeleteUserToken(tokenStr, w)
	}
}

func GetUserDetails(userManager manager.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		ctx := r.Context()
		email := ctx.Value("email").(string)
		response, err := userManager.GetUserDetails(email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

func UpdatePassword(userManager manager.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		var requestbody interface{}
		buffer, err := ioutil.ReadAll(r.Body)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		r.Body.Close()
		json.Unmarshal(buffer, &requestbody)
		body := requestbody.(map[string]interface{})

		ctx := r.Context()
		email := ctx.Value("email").(string)
		response, err := userManager.UpdatePassword(email, body)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

func UpdateProfile(userManager manager.UserManager) http.HandlerFunc {
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
		response, err := userManager.UpdateProfile(email, body)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

func GetAllUsers(userManager manager.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		ctx := r.Context()
		email := ctx.Value("email").(string)
		role := ctx.Value("role").(string)

		err := userManager.AuthorizeUser(role, email)
		if err != nil {
			utils.GetError(err, w)
			return
		}

		users, err := userManager.GetAllUsers(role, email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(users)
	}
}

func GetUser(userManager manager.UserManager) http.HandlerFunc {
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
		role := ctx.Value("role").(string)

		err = userManager.AuthorizeUser(role, email)
		if err != nil {
			utils.GetError(err, w)
			return
		}

		user, err := userManager.GetUser(role, email, id)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}

func UpdateRole(userManager manager.UserManager) http.HandlerFunc {
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
		role := "admin"
		err = userManager.AuthorizeUser(role, email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		response, err := userManager.UpdateRole(body)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}
