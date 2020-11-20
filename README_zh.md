# Cache

<!--构建等icon-->


<!--背景-->


## 安装

`go get github.com/AImager/gocache`

## 使用

```go
import "github.com/AImager/gocache/config"
import "github.com/go-redis/redis/v8"

func main() {
    a, b := 1, 3
	c := &Cache{client: redis.NewClient(&redis.Options{
		Addr: ":6379",
    })}

    // 装饰db查询逻辑，自动缓存
	decoratedGetFunc := getDb
	c.CacheWithFunc(context.TODO(), config.CacheConfig{
		Key:    fmt.Sprintf("test_cache:a:%d:b:%d", a, b),
		Expire: 300,
	}, &decoratedGetFunc, getDb)
    decoratedGetFunc(context.TODO(), a, b)

    // 装饰db更新逻辑，自动删除缓存
	decoratedDelFunc := updateDb
	c.CacheDelWithFunc(context.TODO(), config.CacheDelConfig{
		Key: fmt.Sprintf("test_cache:a:%d:b:%d", a, b),
	}, &decoratedDelFunc, updateDb)
	decoratedDelFunc(context.TODO(), a, b)
}

func getDb(ctx context.Context, a int, b string) (c int, err error) {
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