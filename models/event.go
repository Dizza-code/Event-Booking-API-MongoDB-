package models

import (
	"context"
	"fmt"
	"time"

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

func GetAllEvents() {

}
