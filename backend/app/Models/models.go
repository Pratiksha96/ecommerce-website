package models

//import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Price       int    `json:"price" bson:"price"`
	Ratings     int    `json:"ratings" bson:"ratings"`
	Category    string `json:"category" bson:"category"`
	Stock       string `json:"stock" bson:"stock"`
}
