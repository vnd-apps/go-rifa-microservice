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
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
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
	defer cancel()

	//Connect to the DB
	log.Print(url)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))

	//Check for any errors
	if err != nil {
		log.Fatalf("unable to initialize connection %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("unable to connect %v", err)
	}
	//Print confirmation of connection
	log.Printf("Connected to MongoDB!")
	mdb.Database = client.Database("db-rifa")
	return mdb, nil
}
