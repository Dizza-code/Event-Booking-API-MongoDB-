package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    string `binding:"required"`
	UserID      int64
}

// save function that will save event to the database
var events = []Event{}

func (e Event) Save(collection *mongo.Collection) error {
	// adds it to the database
	eventDoc := map[string]interface{}{
		"name":        e.Name,
		"description": e.Description,
		"location":    e.Location,
		"dateTime":    e.DateTime,
		"user_id":     e.UserID,
	}

	//insert document into the collection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, eventDoc)
	if err != nil {
		return fmt.Errorf("failed to insert event: %w", err)
	}
	return nil
}

func GetAllEvents(collection *mongo.Collection) ([]Event, error) {
	// Create a context with a 10-second timeout
	// This ensures the database operation doesn't hang indefinitely
	// defer cancel() releases resources when the function returns
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// find all documents in the collection
	// Find() without any filter (bson.M{} is an empty filter) returns all documents
	// cursor allows us to iterate through the results efficiently
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch events: %w", err)
	}
	// Ensure cursor is closed after function returns to prevent memory leaks
	defer cursor.Close(ctx)

	// decode all document into Event slice // Initialize slice to store all events
	var events []Event
	// cursor.All decodes all documents into the events slice at once
	if err = cursor.All(ctx, &events); err != nil {
		return nil, fmt.Errorf("failed to decode events: %w", err)
	}
	return events, nil
}
