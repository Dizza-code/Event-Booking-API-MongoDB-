package routes

import (
	"net/http"

	"example.com/events-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

//	func getEvents(context *gin.Context) {
//		events := models.GetAllEvents()
//		context.JSON(http.StatusOK, events)
//	}
//
// Create a handler struct to hold the MongoDB client
type EventHandler struct {
	client *mongo.Client
}

// Create a constructor for the handler
func NewEventHandler(client *mongo.Client) *EventHandler {
	return &EventHandler{
		client: client,
	}
}
func (h *EventHandler) CreateEvent(context *gin.Context) {
	var event models.Event

	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	//get the mongdb collection
	collection := h.client.Database("events_db").Collection("events")

	//Svae the event
	if err := event.Save(collection); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Successfully created", "event": event})

}

func (h *EventHandler) GetEvents(context *gin.Context) {
	// Get reference to the events collection from MongoDB
	collection := h.client.Database("events_db").Collection("events")

	// Call GetAllEvents to retrieve all events from the database
	events, err := models.GetAllEvents(collection)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}
