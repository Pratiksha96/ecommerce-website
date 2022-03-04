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
	Avatar              ProfileImage   `json:"avatar" bson:"avatar"`
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

func (u *User) GetBSON() (interface{}, error) {
	if len(u.Role) == 0 {
		u.Role = "user"
	}
	type my *User
	return my(u), nil
}

func (user *User) GetJwtToken() (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute).Unix()

	claims := &Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(SecretKey))
	return tokenStr, err
}
