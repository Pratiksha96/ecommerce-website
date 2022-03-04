package middleware

import (
	"context"
	"ecommerce-website/app/utils"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func AuthenticateUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookieToken, err := r.Cookie("token")

		if err != nil || len(cookieToken.Value) == 0 {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				utils.GetError(errors.New("Received ErrNoCookie, please login to access this resource"), w)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			utils.GetError(errors.New("Please login to access this resource"), w)
			return
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
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			utils.GetError(errors.New("received error while validating token"), w)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			utils.GetError(errors.New("token is invalid"), w)
			return
		}

		tknClaims, _ := tkn.Claims.(jwt.MapClaims)
		email, _ := tknClaims["email"].(string)
		ctx := r.Context()
		r = r.WithContext(context.WithValue(ctx, "email", email))
		next.ServeHTTP(w, r)
		return
	}
}
