package main

import (
	"log"
	"os"
	"wehrmacht-encyclopedia/database"
	"wehrmacht-encyclopedia/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Get configuration from environment variables
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./wehrmacht.db"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize database
	err := database.InitDB(dbPath)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer database.Close()

	// Seed database with initial data
	if err := database.SeedData(); err != nil {
		log.Fatal("Failed to seed database:", err)
	}

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// API Routes
	api := e.Group("/api")

	// Generals
	api.GET("/generals", handlers.GetGenerals)
	api.GET("/generals/:id", handlers.GetGeneral)

	// Terms
	api.GET("/terms", handlers.GetTerms)

	// Battles
	api.GET("/battles", handlers.GetBattles)
	api.GET("/battles/:id", handlers.GetBattle)

	// Quotes
	api.GET("/quotes/daily", handlers.GetDailyQuote)
	api.GET("/quotes", handlers.GetQuotes)

	// Search
	api.GET("/search", handlers.SearchAll)

	// Serve static assets (CSS, JS, images)
	// IMPORTANT: This must come BEFORE the catch-all route
	e.Static("/assets", "frontend/dist/assets")
	e.File("/vite.svg", "frontend/dist/vite.svg")

	// Serve index.html for root
	e.GET("/", func(c echo.Context) error {
		return c.File("frontend/dist/index.html")
	})

	// Catch-all route for SPA routing (must be LAST)
	// This handles all other routes and serves index.html
	e.GET("/*", func(c echo.Context) error {
		// Don't serve index.html for /api routes (already handled above)
		// Don't serve index.html for /assets routes (already handled above)
		return c.File("frontend/dist/index.html")
	})

	// Start server
	log.Printf("Server starting on :%s", port)
	log.Printf("Database: %s", dbPath)
	e.Logger.Fatal(e.Start(":" + port))
}
