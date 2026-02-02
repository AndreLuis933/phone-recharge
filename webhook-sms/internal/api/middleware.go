package api

import (
	"fmt"
	"httpsms-webhook/internal/logger"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(signingKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			logger.Log.Error("Requisição sem header Authorization")
			c.JSON(401, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			logger.Log.Error("Header Authorization em formato inválido (faltou 'Bearer')")
			c.JSON(401, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("algoritmo inesperado: %v", token.Header["alg"])
			}
			return []byte(signingKey), nil
		})

		if err != nil {
			logger.Log.Error("Erro ao validar JWT", "error", err)
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		if !token.Valid {
			logger.Log.Error("Token inválido")
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		logger.Log.Info("JWT validado com sucesso")
		c.Next()
	}
}