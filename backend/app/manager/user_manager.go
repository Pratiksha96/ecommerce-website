package manager

import (
	"context"
	models "ecommerce-website/app/models"
	"ecommerce-website/app/utils"
	"ecommerce-website/internal/database"
	"log"

	"golang.org/x/crypto/bcrypt"

	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserManager interface {
	RegisterUser(user models.User) (TokenResponse, error)
	LoginUser(user models.User) (TokenResponse, error)
	GetUserDetails(email string) (*models.User, error)
	UpdatePassword(email string, body map[string]interface{}) (UserResponse, error)
	UpdateProfile(email string, body map[string]interface{}) (UserResponse, error)
	GetAllUsers(role string, email string) ([]primitive.M, error)
	GetUser(role string, email string, id primitive.ObjectID) (*models.User, error)
}

type TokenResponse struct {
	Success bool
	Token   string
}

type UserResponse struct {
	Success bool
	User    models.User
}

type userManager struct{}

func NewUserManager() UserManager {
	return &userManager{}
}

func (um *userManager) RegisterUser(user models.User) (TokenResponse, error) {

	user.GetBSON()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return TokenResponse{}, err
	}
	user.Password = string(hashedPassword)

	registeredUser, err := database.Coll_user.InsertOne(context.TODO(), user)
	if err != nil {
		return TokenResponse{}, err
	}

	token, err := user.GetJwtToken()
	if err != nil {
		return TokenResponse{}, err
	}
	log.Println("User registered successfully with id: ", user.Email, registeredUser.InsertedID)

	var tokenResponse = TokenResponse{
		Success: true,
		Token:   token,
	}
	return tokenResponse, nil
}

func (um *userManager) LoginUser(user models.User) (TokenResponse, error) {
	var storedUser models.User
	filter := bson.M{"email": user.Email}
	err := database.Coll_user.FindOne(context.TODO(), filter).Decode(&storedUser)

	if err != nil {
		return TokenResponse{}, errors.New("No such user present")
	}

	passwordErr := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if passwordErr != nil {
		return TokenResponse{}, errors.New("Password mismatched")
	}

	token, err := user.GetJwtToken()
	if err != nil {
		return TokenResponse{}, err
	}

	log.Println("User logged in successfully!")
	var tokenResponse = TokenResponse{
		Success: true,
		Token:   token,
	}
	return tokenResponse, nil
}

func (um *userManager) GetUserDetails(email string) (*models.User, error) {
	var storedUser models.User
	filter := bson.M{"email": email}
	err := database.Coll_user.FindOne(context.TODO(), filter).Decode(&storedUser)
	if err != nil {
		return nil, errors.New("no such user present")
	}
	return &storedUser, nil
}

func (um *userManager) UpdatePassword(email string, body map[string]interface{}) (UserResponse, error) {
	var storedUser models.User
	filter := bson.M{"email": email}
	err := database.Coll_user.FindOne(context.TODO(), filter).Decode(&storedUser)
	if err != nil {
		return UserResponse{}, errors.New("No such user present")
	}

	oldPassword := body["oldPassword"].(string)
	newPassword := body["newPassword"].(string)
	confirmPassword := body["confirmPassword"].(string)
	passwordErr := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(oldPassword))
	if passwordErr != nil {
		return UserResponse{}, errors.New("password mismatched")
	}
	if confirmPassword != newPassword {
		return UserResponse{}, errors.New("new password do not match with confirm password")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return UserResponse{}, err
	}

	result, err := database.Coll_user.UpdateOne(
		context.TODO(),
		bson.M{"email": email},
		bson.D{
			{"$set", bson.D{{"password", string(hashedPassword)}}},
		},
	)
	if err != nil {
		return UserResponse{}, errors.New("Failed to update password")
	}
	storedUser.Password = string(hashedPassword)
	log.Println("Following number of users updated ", result.ModifiedCount)
	var response = UserResponse{
		Success: true,
		User:    storedUser,
	}
	return response, nil
}

func (um *userManager) UpdateProfile(email string, body map[string]interface{}) (UserResponse, error) {
	var storedUser models.User
	filter := bson.M{"email": email}
	err := database.Coll_user.FindOne(context.TODO(), filter).Decode(&storedUser)
	if err != nil {
		return UserResponse{}, errors.New("no such user present")
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
			return UserResponse{}, errors.New("Failed to update profile information")
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
			return UserResponse{}, errors.New("Failed to update profile information")
		}
		log.Println("Following number of users updated ", result.ModifiedCount)
	}
	storedUser.Name = newName
	storedUser.Email = newEmail

	var response = UserResponse{
		Success: true,
		User:    storedUser,
	}
	return response, nil
}

func (um *userManager) GetAllUsers(role string, email string) ([]primitive.M, error) {

	cur, err := database.Coll_user.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			return nil, e
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	} else if len(results) == 0 {
		return nil, errors.New("User list is empty")
	}

	cur.Close(context.Background())
	payload := results
	return payload, nil
}

func (um *userManager) GetUser(role string, email string, id primitive.ObjectID) (*models.User, error) {
	err := utils.AuthorizeUser(role, email)
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	filter := bson.M{"_id": id}
	err = database.Coll_user.FindOne(context.TODO(), filter).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
