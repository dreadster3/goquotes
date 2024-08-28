package main

import (
	"github.com/dreadster3/goquotes/api"
	"github.com/dreadster3/goquotes/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	err := database.InitDB()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	apiGroup := router.Group("/api")
	api.RegisterRoutes(apiGroup)

	router.Run(":8080")
}
