# 滑动窗口限流器

[English](README.md)

## 概览

`Sliding Window Limiter` 是一个轻量、易用的 Go 库，用于通过滑动窗口算法实现限流。它使用 Redis 提供分布式限流功能。

## 特性

- 简单集成 `redis/v8`
- 使用滑动窗口算法实现精确限流
- 支持 Redis 连接的简单配置
- 可扩展到分布式系统

## 安装

```bash
go get github.com/sktli/slidingwindow
```

## 使用示例

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
            fmt.Printf("请求 %d 被允许
", i+1)
        } else {
            fmt.Printf("请求 %d 被限流
", i+1)
        }
        time.Sleep(500 * time.Millisecond)
    }
}
```

## 开源协议

本项目采用 MIT 协议，详细内容请参阅 [LICENSE](LICENSE) 文件。
