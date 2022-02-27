//Creating token and saving it in Cookie
package utils

import (
	"net/http"
	"time"
)

func StoreUserToken(token string, w http.ResponseWriter) {

	cookie := &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().AddDate(0, 0, 7),
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(200)
}
