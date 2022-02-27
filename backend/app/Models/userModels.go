package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

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

	// Create the JWT claims, which includes the user email and expiry time
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Email,
		ExpiresAt: expirationTime,
	})

	token, err := claims.SignedString([]byte(SecretKey))
	return token, err
}
