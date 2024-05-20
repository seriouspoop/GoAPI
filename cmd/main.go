package main

import (
	"context"
	"log"

	"github.com/seriouspoop/GoAPI/cmd/api"
	"github.com/seriouspoop/GoAPI/config"
	"github.com/seriouspoop/GoAPI/db"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {

	db, err := db.NewMongoStorage(config.Envs)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Client().Disconnect(context.TODO())

	initStorage(db)

	server := api.NewAPIServer(":3000", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *mongo.Database) {
	err := db.Client().Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB connection successfull!!")
}
