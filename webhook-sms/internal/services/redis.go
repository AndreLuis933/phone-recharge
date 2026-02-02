// services/redis.go
package services

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	ctx         = context.Background()
)

func InitRedis(addr, password string) error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	return redisClient.Ping(ctx).Err()
}

func SaveVivoCode(code string) error {

	// Salva com TTL de 5 minutos
	return redisClient.Set(ctx, "vivo:code", code, 5*time.Minute).Err()
}
