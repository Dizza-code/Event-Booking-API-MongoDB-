package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var Client *mongo.Client

func InitDB(uri string) {

	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	//set up context for connection
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	// defer func() {
	// 	if err := client.Disconnect(ctx); err != nil {
	// 		log.Fatalf("Failed to disconnect MongoDB client: %v", err)
	// 	}
	// }()

	// Ping the database to verify connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB successfully!")
	Client = client

	_ = client.Ping(ctx, readpref.Primary())
}

// GetCollection gets a specific collection from the database
func GetCollection(databaseName, collectionName string) *mongo.Collection {
	return Client.Database(databaseName).Collection(collectionName)
}

// CloseConnection closes the connection to MongoDB
func CloseConnection() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := Client.Disconnect(ctx)
	if err != nil {
		log.Fatalf("Failed to disconnect MongoDB client: %v", err)
	}
}

func createCollections(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//create events collection
	err := client.Database("events_db").CreateCollection(ctx, "events")
	if err != nil {
		//check for collection already exists
		if cmdErr, ok := err.(mongo.CommandError); ok && cmdErr.Code == 48 {
			log.Println("Events collection already exists")
			return nil
		}
		return fmt.Errorf("failed to create events collection: %w", err)
	}
	return nil
}
