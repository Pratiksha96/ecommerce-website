package manager

import (
	"context"
	models "ecommerce-website/app/models"
	"ecommerce-website/app/utils"
	"ecommerce-website/internal/database"
	"encoding/json"
	"log"

	"golang.org/x/crypto/bcrypt"

	"net/http"

	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterUser(user models.User, w http.ResponseWriter) {

	user.GetBSON()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)

	registeredUser, err := database.Coll_user.InsertOne(context.TODO(), user)
	if err != nil {
		utils.GetError(err, w)
		return
	}

	token, err := user.GetJwtToken()
	if err != nil {
		utils.GetError(err, w)
		return
	}
	log.Println("User registered successfully with id: ", user.Email, registeredUser.InsertedID)
	utils.StoreUserToken(token, w)
	tokenResponse := map[string]interface{}{"success": true, "token": token}
	json.NewEncoder(w).Encode(tokenResponse)
}

func LoginUser(user models.User, w http.ResponseWriter) {

	var storedUser models.User
	filter := bson.M{"email": user.Email}
	err := database.Coll_user.FindOne(context.TODO(), filter).Decode(&storedUser)

	if err != nil {
		utils.GetError(errors.New("no such user present"), w)
		return
	}

	passwordErr := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if passwordErr != nil {
		utils.GetError(errors.New("password mismatched"), w)
		return
	}

	token, err := user.GetJwtToken()
	if err != nil {
		utils.GetError(err, w)
		return
	}

	log.Println("User logged in successfully!")
	utils.StoreUserToken(token, w)

	tokenResponse := map[string]interface{}{"success": true, "token": token}
	json.NewEncoder(w).Encode(tokenResponse)
}

func LogoutUser(token string, w http.ResponseWriter) {
	utils.DeleteUserToken(token, w)
}

func GetUserDetails(email string, w http.ResponseWriter) {
	var storedUser models.User
	filter := bson.M{"email": email}
	err := database.Coll_user.FindOne(context.TODO(), filter).Decode(&storedUser)
	if err != nil {
		utils.GetError(errors.New("no such user present"), w)
		return
	}
	json.NewEncoder(w).Encode(storedUser)

}

func UpdatePassword(email string, body map[string]interface{}, w http.ResponseWriter) {

	var storedUser models.User
	filter := bson.M{"email": email}
	err := database.Coll_user.FindOne(context.TODO(), filter).Decode(&storedUser)
	if err != nil {
		utils.GetError(errors.New("no such user present"), w)
		return
	}

	oldPassword := body["oldPassword"].(string)
	newPassword := body["newPassword"].(string)
	confirmPassword := body["confirmPassword"].(string)
	passwordErr := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(oldPassword))
	if passwordErr != nil {
		utils.GetError(errors.New("password mismatched"), w)
		return
	}
	if confirmPassword != newPassword {
		utils.GetError(errors.New("new password do not match with confirm password"), w)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	result, err := database.Coll_user.UpdateOne(
		context.TODO(),
		bson.M{"email": email},
		bson.D{
			{"$set", bson.D{{"password", string(hashedPassword)}}},
		},
	)
	if err != nil {
		utils.GetError(errors.New("Failed to update password"), w)
		return
	}
	storedUser.Password = string(hashedPassword)
	log.Println("Following number of users updated ", result.ModifiedCount)
	response := map[string]interface{}{"success": true, "user": storedUser}
	json.NewEncoder(w).Encode(response)

}

func UpdateProfile(email string, body map[string]interface{}, w http.ResponseWriter) {

	var storedUser models.User
	filter := bson.M{"email": email}
	err := database.Coll_user.FindOne(context.TODO(), filter).Decode(&storedUser)
	if err != nil {
		utils.GetError(errors.New("no such user present"), w)
		return
	}

	newName := body["name"].(string)
	newEmail := body["email"].(string)
	if newEmail != storedUser.Email {
		result, err := database.Coll_user.UpdateOne(
			context.TODO(),
			bson.M{"email": email},
			bson.D{
				{"$set", bson.D{{"name", newName}, {"email", newEmail}}},
			},
		)
		if err != nil {
			utils.GetError(errors.New("Failed to update profile information"), w)
			return
		}
		log.Println("Following number of users updated ", result.ModifiedCount)
	} else {
		result, err := database.Coll_user.UpdateOne(
			context.TODO(),
			bson.M{"email": email},
			bson.D{
				{"$set", bson.D{{"name", newName}}},
			},
		)
		if err != nil {
			utils.GetError(errors.New("Failed to update profile information"), w)
			return
		}
		log.Println("Following number of users updated ", result.ModifiedCount)
	}
	storedUser.Name = newName
	storedUser.Email = newEmail

	response := map[string]interface{}{"success": true, "user": storedUser}
	json.NewEncoder(w).Encode(response)

}

func GetAllUsers(role string, email string, w http.ResponseWriter) {

	err := utils.AuthorizeUser(role, email)
	if err != nil {
		utils.GetError(err, w)
		return
	}

	cur, err := database.Coll_user.Find(context.Background(), bson.D{{}})
	if err != nil {
		utils.GetError(err, w)
		return
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			utils.GetError(err, w)
			return
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		utils.GetError(err, w)
		return
	} else if len(results) == 0 {
		utils.GetError(errors.New("User list is empty"), w)
		return
	}

	cur.Close(context.Background())
	payload := results
	json.NewEncoder(w).Encode(payload)

}

func GetSingleUser(role string, email string, id primitive.ObjectID, w http.ResponseWriter) {

	err := utils.AuthorizeUser(role, email)
	if err != nil {
		utils.GetError(err, w)
		return
	}

	user := &models.User{}
	filter := bson.M{"_id": id}
	err = database.Coll_user.FindOne(context.TODO(), filter).Decode(user)
	if err != nil {
		utils.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(user)

}
