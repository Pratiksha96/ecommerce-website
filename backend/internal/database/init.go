package database

import (
	"context"
	"fmt"
	"time"

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
	fmt.Println("Connection Established with database")
}
