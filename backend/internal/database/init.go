package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db           *mongo.Database
	Coll_product *mongo.Collection
	Coll_user    *mongo.Collection
)

func InitDB() {
	uri := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	db := client.Database("ecommerce_website")
	Coll_product = db.Collection("product")

	Coll_user = db.Collection("user")

	indexName, err := Coll_user.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Email Index created in user collection: ", indexName)
	log.Println("Connection Established with database")
}
