package db

import (
	"context"
	"log"

	"github.com/seriouspoop/GoAPI/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoStorage(cfg config.Config) (db *mongo.Database, err error) {
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(cfg.FormatURI()))
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database(cfg.Database.DB)
	return db, nil
}
