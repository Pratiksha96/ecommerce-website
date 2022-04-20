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

//Handler that entertains Register User API
func RegisterUser(userManager manager.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		//Getting user details from request body
		var user models.User
		_ = json.NewDecoder(r.Body).Decode(&user)

		//validating user details
		if errs := utils.UserRegisterValidation(user); len(errs) > 0 {
			utils.GetErrorWithStatus(errors.New("Received invalid json request!"), w, http.StatusBadRequest)
			return
		}
		//Registering user and creating token for newly created user
		tokenResponse, err := userManager.RegisterUser(user)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		//Storing user token just created by registering user
		utils.StoreUserToken(tokenResponse.Token, w)
		json.NewEncoder(w).Encode(tokenResponse)
	}
}

//Handler that entertains Login User API
func LoginUser(userManager manager.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		//Decoding login details of user
		var user models.User
		_ = json.NewDecoder(r.Body).Decode(&user)

		//Validating user details
		if errs := utils.UserLoginValidation(user); len(errs) > 0 {
			utils.GetErrorWithStatus(errors.New("Received invalid json request!"), w, http.StatusBadRequest)
			return
		}
		//Calling function to validate valid username password and then creating access tokem
		tokenResponse, err := userManager.LoginUser(user)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		utils.StoreUserToken(tokenResponse.Token, w)
		json.NewEncoder(w).Encode(tokenResponse)
	}
}

//Handler that will entertain logout user API
func LogoutUser(userManager manager.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		//fetching token from the request
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
		//Deleting token that was earlier set for particular user
		tokenStr := cookieToken.Value
		utils.DeleteUserToken(tokenStr, w)
	}
}

//Handler that entertains request to get user details
func GetUserDetails(userManager manager.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		ctx := r.Context()
		email := ctx.Value("email").(string)
		// Calling user manager to get user details
		response, err := userManager.GetUserDetails(email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

//Handeler that entertains Update password requests
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
		//Getting details from request body
		body := requestbody.(map[string]interface{})

		ctx := r.Context()
		email := ctx.Value("email").(string)
		//Calling userMAnager to update password
		response, err := userManager.UpdatePassword(email, body)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

//Handler that entertains update profile information API request
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
		//fetch user details from request body
		body := requestbody.(map[string]interface{})

		ctx := r.Context()
		email := ctx.Value("email").(string)
		//Call user manager to update profile information
		response, err := userManager.UpdateProfile(email, body)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

//Handler that entertains get allusers API request
func GetAllUsers(userManager manager.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		ctx := r.Context()
		email := ctx.Value("email").(string)
		role := ctx.Value("role").(string)
		// Authorizing user, Only admin can access this resource
		err := userManager.AuthorizeUser(role, email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		//call userManager to get all users information
		users, err := userManager.GetAllUsers(role, email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(users)
	}
}

//Handler that entertains user details request for particular user
func GetUser(userManager manager.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//getting user id from URL parameters
		var params = mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil {
			utils.GetErrorWithStatus(errors.New("invalid object id"), w, http.StatusUnprocessableEntity)
			return
		}
		ctx := r.Context()
		email := ctx.Value("email").(string)
		role := ctx.Value("role").(string)
		// Authorizing user, Only admin can access this resource
		err = userManager.AuthorizeUser(role, email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		//Calling user manager to get particular user
		user, err := userManager.GetUser(role, email, id)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}

//Handler that entertains update role API request
func UpdateRole(userManager manager.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		//get user role details from request body
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
		// Authorizing user, Only admin can access this resource
		err = userManager.AuthorizeUser(role, email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		//Calling user manager to update user role(admin/user)
		response, err := userManager.UpdateRole(body)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

//Handler that entertains delete user request
func DeleteUser(userManager manager.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		//get user id to be deleted from URL parameters
		var params = mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil || params["id"] == "" {
			utils.GetErrorWithStatus(errors.New("invalid object id"), w, http.StatusUnprocessableEntity)
			return
		}
		ctx := r.Context()
		email := ctx.Value("email").(string)
		deleteResponse, err := userManager.DeleteUser(id, "admin", email)
		if err != nil {
			utils.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(deleteResponse)
	}
}
