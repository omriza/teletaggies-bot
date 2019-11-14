package data

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitRepository() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	clientOptions.Auth.Username = "root"
	clientOptions.Auth.Password = "p$$wd"
	//clientOptions.Auth.AuthMechanism = "PLAIN"

	// Connect to MongoDB
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("Connected to MongoDB!")
}
