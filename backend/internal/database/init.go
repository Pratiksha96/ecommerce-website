package database

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Coll_product *mongo.Collection
	Coll_user    *mongo.Collection
	Coll_order   *mongo.Collection
)

func InitDB() {
	uri := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	db := client.Database("ecommerce_website")
	Coll_product = db.Collection("product")

	Coll_user = db.Collection("user")

	Coll_order = db.Collection("order")

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

	log.Debug("Email Index created in user collection: ", indexName)
	log.Info("Connection Established with database")
}
