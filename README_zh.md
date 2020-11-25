# GoCache

🇬🇧 [English](README.md) | 🇨🇳 中文

## 安装

`go get github.com/AImager/gocache`

## 使用

```go
package main

import (
	"context"
	"fmt"

	cache "github.com/AImager/gocache"
	"github.com/AImager/gocache/config"
)

func main() {
	a, b := 1, 3
	client, _ := cache.GetClient(config.ClientConfig{
		Addr:       "127.0.0.1:6379",
		ClientType: config.Goredis,
	})
	c := &cache.Cache{Client: client}

	// 装饰获取方法，自动设置缓存
	decoratedGetFunc := getDb
	c.CacheWithFunc(context.TODO(), config.CacheConfig{
		Key:    fmt.Sprintf("test_cache:a:%d:b:%d", a, b),
		Expire: 300,
	}, &decoratedGetFunc, getDb)
	decoratedGetFunc(context.TODO(), a, b)

	// 装饰更新方法，自动删除缓存
	decoratedUpdateFunc := updateDb
	c.CacheDelWithFunc(context.TODO(), config.CacheDelConfig{
		Key: fmt.Sprintf("test_cache:a:%d:b:%d", a, b),
	}, &decoratedUpdateFunc, updateDb)
	decoratedUpdateFunc(context.TODO(), a, b)
}

func getDb(ctx context.Context, a int, b int) (c int, err error) {
	return 1, nil
}

func updateDb(ctx context.Context, a int, b int) (err error) {
	return nil
}
```

## 贡献者

PRs accepted.

## 许可

MIT © AImager.