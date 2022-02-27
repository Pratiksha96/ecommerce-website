package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//TODO to be changed
const SecretKey = "ThisIsMySecretKey"

type User struct {
	Name                string          `json:"name" bson:"name"`
	Email               string          `json:"email" bson:"email"`
	Password            string          `json:"password" bson:"password"`
	Avatar              []*ProfileImage `json:"avatar" bson:"avatar"`
	Role                string          `json:"role" bson:"role"`
	ResetPasswordToken  string          `json:"resetPasswordToken" bson:"resetPasswordToken"`
	ResetPasswordExpire time.Time       `json:"resetPasswordExpire" bson:"resetPasswordExpire"`
}

type ProfileImage struct {
	Public_id string `json:"public_id" bson:"public_id"`
	Url       string `json:"url" bson:"url"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

//setting role as user by default
func (u *User) GetBSON() (interface{}, error) {
	if len(u.Role) == 0 {
		u.Role = "user"
	}
	type my *User
	return my(u), nil
}

//get user token
func (user *User) GetJwtToken() (string, error) {

	// Declaring the expiration time of the user  token here
	// We have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute).Unix()

	claims := &Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime,
		},
	}

	// Create the JWT claims, which includes the user email and expiry time
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(SecretKey))
	return tokenStr, err
}
