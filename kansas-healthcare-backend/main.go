package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"kansas-healthcare-api/config"
	"kansas-healthcare-api/controllers"
	"kansas-healthcare-api/data"
	"kansas-healthcare-api/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize data repository
	var repo data.Repository
	if cfg.DataSource == "json" {
		repo = data.NewJSONRepository()
	} else {
		// Future: repo = data.NewDBRepository()
		log.Fatal("Database repository not implemented yet")
	}

	// Initialize services
	providerService := services.NewProviderService(repo)
	analyticsService := services.NewAnalyticsService(repo)

	// Initialize controllers
	providerController := controllers.NewProviderController(providerService)
	analyticsController := controllers.NewAnalyticsController(analyticsService)

	// Setup router
	r := gin.Default()

	// CORS middleware
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5173"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(corsConfig))

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"timestamp": time.Now().UTC(),
			"service": "kansas-healthcare-api",
		})
	})

	// Routes
	api := r.Group("/api/v1")
	{
		api.GET("/providers", providerController.GetProviders)
		api.GET("/provider-network", providerController.GetProviderNetwork)
		api.GET("/county-data/:county", analyticsController.GetCountyData)
		api.GET("/county-data", analyticsController.GetAllCountyData)
		api.GET("/recommendations/:county", analyticsController.GetRecommendations)
		api.POST("/filters", providerController.GetFilteredData)
		api.GET("/active-providers", analyticsController.GetActiveProviderCount)
		api.GET("/terminated-analysis", analyticsController.GetTerminatedNetworkAnalysis)
		api.GET("/terminated-analysis/:county", analyticsController.GetCountyTerminatedNetworkAnalysis)
		api.GET("/specialty-density/:county", analyticsController.GetSpecialtyDensityAnalysis)
		api.GET("/radius-analysis/:county", analyticsController.GetRadiusAnalysis)
	}

	port := cfg.Port
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on port %s using %s data source", port, cfg.DataSource)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Graceful shutdown with 30 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}
