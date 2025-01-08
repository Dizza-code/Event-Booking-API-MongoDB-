package routes

import (
	"example.com/events-api/db"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Create handler with the MongoDB client
	eventHandler := NewEventHandler(db.Client)
	server.GET("events", eventHandler.GetEvents)
	server.POST("events", eventHandler.CreateEvent)
}
