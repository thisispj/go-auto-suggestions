package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thisispj/autosuggestions/cache"
	"github.com/thisispj/autosuggestions/handler"
	"github.com/thisispj/autosuggestions/service"
	"github.com/thisispj/autosuggestions/store"
)

func main() {
	// Initialize repository
	repo, err := store.NewCountryStore("assets/data.json")
	if err != nil {
		panic(err)
	}

	// Initialize Redis cache
	redisCache := cache.NewRedisCache()

	// Initialize service
	service := service.NewCountryService(redisCache, repo)

	// Initialize handler
	handler := handler.NewCountryHandler(service)

	// Initialize Gin router
	router := gin.Default()

	// Define routes
	router.GET("/autocomplete", handler.AutocompleteHandler)

	// Serve HTML frontend
	router.StaticFile("/", "./static/index.html")

	// Start server
	router.Run(":8080")
}
