package storage

import (
	"testing"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestRedisConnection(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", 
		DB:       0,  
	})

	pong, err := client.Ping(client.Context()).Result()

	assert.NoError(t, err)
	assert.Equal(t, "PONG", pong)
}