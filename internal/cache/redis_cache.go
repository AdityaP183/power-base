package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisCache implements caching using Redis
type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

// NewRedisCache creates a new Redis cache instance
func NewRedisCache(addr, password string, db int) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisCache{
		client: client,
		ctx:    context.Background(),
	}
}

// Get retrieves a value from cache
func (c *RedisCache) Get(key string) ([]byte, error) {
	val, err := c.client.Get(c.ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	return val, nil
}

// Set stores a value in cache with TTL
func (c *RedisCache) Set(key string, value []byte, ttl int) error {
	return c.client.Set(c.ctx, key, value, time.Duration(ttl)*time.Second).Err()
}

// Delete removes a value from cache
func (c *RedisCache) Delete(key string) error {
	return c.client.Del(c.ctx, key).Err()
}

// Close closes the Redis connection
func (c *RedisCache) Close() error {
	return c.client.Close()
}
