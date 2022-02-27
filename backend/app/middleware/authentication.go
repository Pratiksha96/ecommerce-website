package middleware

import (
	"ecommerce-website/app/utils"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func AuthenticateUser(w http.ResponseWriter, r *http.Request) (bool, string) {

	cookieToken, err := r.Cookie("token")

	if err != nil || len(cookieToken.Value) == 0 {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			utils.GetError(errors.New("received ErrNoCookie, please login to access this resource"), w)
			return false, ""
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		utils.GetError(errors.New("please login to access this resource"), w)
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
