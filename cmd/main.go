package main

import (
	"github.com/gofiber/fiber/v2"
	config "github.com/ipekutku/playlist-generator"
	auth "github.com/ipekutku/playlist-generator/pkg/auth"
	log "github.com/ipekutku/playlist-generator/pkg/log"
)

func main() {
	// Initialize logger
	logger := log.NewLogger()

	// Defer Sync to ensure all logs are flushed before the program exits
	defer logger.Sync()

	// Load configuration
	appConfig := config.SpotifyConfig()

	// Initialize Spotify service and AuthHandler
	spotify := auth.NewSpotifyService(&appConfig.SpotifyLoginConfig)
	auth := auth.NewAuthHandler(spotify, logger, &appConfig.SpotifyLoginConfig)

	// Initialize Fiber app
	app := fiber.New()

	// Define routes
	app.Get("/login", auth.SpotifyLogin)
	app.Post("/callback", auth.SpotifyCallback)

	// Start server
	if err := app.Listen(":8080"); err != nil {
		logger.Error("Failed to start server", "error", err)
	}
}
