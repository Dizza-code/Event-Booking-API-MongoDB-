package main

import (
	"example.com/events-api/db"
	"example.com/events-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB("mongodb://localhost:27017")

	defer db.CloseConnection()
	server := gin.Default()

	routes.RegisterRoutes(server)
	server.Run(":8080")
}
