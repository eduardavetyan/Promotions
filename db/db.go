package db

import (
	"context"
	"log"
	"promotions-app/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func Init() {
	c := config.GetConfig()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(c.GetString("db.uri")))

	if err != nil {
		log.Fatal(err)
	}

	db = client.Database(c.GetString("db.name"))
}

func GetDB(coll string) *mongo.Collection {
	return db.Collection(coll)
}
