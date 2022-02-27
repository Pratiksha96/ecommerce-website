//Creating token and saving it in Cookie

package utils

import (
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
)

var hashKey = []byte("very-secret")
var s = securecookie.New(hashKey, nil)

func StoreUserToken(token string, w http.ResponseWriter) {

	encoded, err := s.Encode("jwt-token", token)

	if err != nil {
		GetError(err, w)
		return
	}

	cookie := &http.Cookie{
		Name:     "jwt-token", // <- should be any unique key you want
		Value:    encoded,     // <- the token after encoded by SecureCookie
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		Expires:  time.Now().AddDate(0, 0, 7), //7 days limit
	}

	http.SetCookie(w, cookie)

}
