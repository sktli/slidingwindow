# Sliding Window Limiter

[中文版本 (Chinese)](README_zh.md)

## Overview

`Sliding Window Limiter` is a lightweight, easy-to-use Go library for implementing rate limiting using the sliding window algorithm. It leverages Redis to provide distributed rate limiting functionality.

## Features

- Simple integration with `redis/v8`
- Sliding window algorithm for precise rate limiting
- Easy configuration for Redis connection
- Scalable for distributed systems

## Installation

```bash
go get github.com/sktli/slidingwindow
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/sktli/slidingwindow"
)

func main() {
    config := slidingwindow.Config{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    }

    limiter := slidingwindow.NewSlidingWindowLimiter(config)

    ctx := context.Background()
    key := "user:123"
    limit := 5
    window := time.Minute

    for i := 0; i < 10; i++ {
        allowed, err := limiter.Allow(ctx, key, limit, window)
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }
        if allowed {
            fmt.Printf("Request %d allowed
", i+1)
        } else {
            fmt.Printf("Request %d rate-limited
", i+1)
        }
        time.Sleep(500 * time.Millisecond)
    }
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
