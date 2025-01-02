package slidingwindow

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

// SlidingWindowLimiter encapsulates the Redis client and sliding window logic.
type SlidingWindowLimiter struct {
	client *redis.Client
}

// Config defines the configuration for the Redis connection.
type Config struct {
	Addr     string // Redis server address
	Password string // Redis password
	DB       int    // Redis database
}

// NewSlidingWindowLimiter creates a new instance of SlidingWindowLimiter.
func NewSlidingWindowLimiter(cfg Config) *SlidingWindowLimiter {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	return &SlidingWindowLimiter{client: client}
}

func (l *SlidingWindowLimiter) Del(ctx context.Context, key string) error {
	return l.client.Del(ctx, key).Err()
}

// Allow determines if the request can proceed based on the sliding window rate limit.
func (l *SlidingWindowLimiter) Allow(ctx context.Context, key string, limit int, window time.Duration) (bool, error) {
	now := time.Now().UnixNano()

	pipe := l.client.TxPipeline()

	// Remove expired requests
	pipe.ZRemRangeByScore(ctx, key, "0", strconv.FormatInt(now-window.Nanoseconds(), 10))

	// Add the current request
	pipe.ZAdd(ctx, key, &redis.Z{
		Score:  float64(now),
		Member: now,
	})

	// Count the requests in the current window
	zCardCmd := pipe.ZCard(ctx, key)

	// Set key expiration
	pipe.Expire(ctx, key, window)

	// Execute pipeline
	_, err := pipe.Exec(ctx)
	if err != nil {
		return false, err
	}

	// Check if the limit is exceeded
	if zCardCmd.Val() > int64(limit) {
		return false, nil
	}

	return true, nil
}
