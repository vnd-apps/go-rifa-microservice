// Package mongodb implements mongodb connection.
package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	_defaultConnTimeout = time.Hour
)

type MongoDB struct {
	Database *mongo.Database
}

type MongoCol struct {
	Collection *mongo.Collection
}

// New -.
func New(url string, args ...interface{}) (*MongoDB, error) {
	mdb := &MongoDB{}
	ctx, cancel := context.WithTimeout(context.Background(), _defaultConnTimeout)

	// Connect to the DB
	log.Print(url)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	// Check for any errors
	if err != nil {
		log.Fatalf("unable to initialize connection %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("unable to connect %v", err)
	}
	// Print confirmation of connection
	log.Printf("Connected to MongoDB!")

	mdb.Database = client.Database("db-rifa")

	defer cancel()

	return mdb, nil
}
