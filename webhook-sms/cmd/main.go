package main

import (
	"httpsms-webhook/internal/api"
	"httpsms-webhook/internal/logger"
	"httpsms-webhook/internal/services"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Init()

	if err := godotenv.Load(); err != nil {
		logger.Log.Info("No .env file found")
	}

	signingKey := os.Getenv("SIGNING_KEY")
	signingKey = strings.Trim(signingKey, "' \t\n\r")
	if signingKey == "" {
		logger.Fatal("SIGNING_KEY environment variable is required")
	}

	if err := services.InitRedis("redis:6379", ""); err != nil {
		logger.Fatalf("❌ Erro ao conectar no Redis: %v", err)
	}

	logger.Log.Info("✅ Redis conectado")

	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/health", api.Health)
	router.POST("/webhook", api.ValidateJWT(signingKey), api.WebhookHandler)

	if err := router.Run(":8080"); err != nil {
		logger.Fatal("❌ Erro ao iniciar servidor", "error", err, "port", 8080)
	}
}
