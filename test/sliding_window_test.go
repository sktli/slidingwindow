package test

import (
	"context"
	"github.com/sktli/slidingwindow"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRateLimit(t *testing.T) {
	config := slidingwindow.Config{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
	limiter := slidingwindow.NewSlidingWindowLimiter(config)
	limiter.Del(context.Background(), "test")
	for i := 0; i <= 11; i++ {
		ok, _ := limiter.Allow(context.Background(), "test", 10, time.Second)
		t.Log(i, ok)
		if i == 10 || i == 11 {
			assert.Equal(t, false, ok)
		} else {
			assert.Equal(t, true, ok)
		}
	}
	limiter.Del(context.Background(), "test")
}
