package main

import (
	"httpsms-webhook/internal/api"
	"httpsms-webhook/internal/utils"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var signingKey string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	signingKey = os.Getenv("SIGNING_KEY")
	signingKey = strings.Trim(signingKey, "' \t\n\r")
	if signingKey == "" {
		panic("SIGNING_KEY environment variable is required")
	}
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/health", api.Health)
	router.POST("/webhook", utils.ValidateJWT(signingKey), api.WebhookHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("❌ Erro ao iniciar servidor:", err)
	}
}
