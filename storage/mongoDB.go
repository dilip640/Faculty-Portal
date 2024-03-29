package storage

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// InitializeMongo start connection to mongo
func InitializeMongo() {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_CONN_URI"))

	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

}
