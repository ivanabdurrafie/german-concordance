package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivanabdurrafie/german-concordance/pkg/api"
	"github.com/ivanabdurrafie/german-concordance/pkg/config"
	"go.uber.org/zap"
)

func main() {
	// Load config
	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		logger, _ := zap.NewProduction()
		logger.Fatal("Failed to load config", zap.Error(err))
	}

	// Initialize logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Setup router
	router := gin.New()
	router.Use(gin.Recovery())

	// Routes
	router.GET("/health", api.HealthCheck)
	router.POST("/concordance", api.ConcordanceHandler)

	// Start server
	logger.Info("Starting server", zap.String("port", cfg.Server.Port))
	if err := http.ListenAndServe(":"+cfg.Server.Port, router); err != nil {
		logger.Fatal("Server failed", zap.Error(err))
	}
}
