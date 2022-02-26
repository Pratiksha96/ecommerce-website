package models

import (
	"time"
)

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
