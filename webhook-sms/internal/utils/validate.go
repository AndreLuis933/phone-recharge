package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


func ValidateJWT(signingKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			log.Println("❌ Requisição sem header Authorization")
			c.JSON(401, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			log.Println("❌ Header Authorization em formato inválido (faltou 'Bearer')")
			c.JSON(401, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("algoritmo inesperado: %v", token.Header["alg"])
			}
			return []byte(signingKey), nil
		})

		if err != nil {
			log.Printf("❌ Erro ao validar JWT: %v\n", err)
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		if !token.Valid {
			log.Println("❌ Token inválido")
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		log.Println("✅ JWT validado com sucesso")
		c.Next()
	}
}