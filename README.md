# Cache

## Install

`go get github.com/AImager/gocache`

## Usage

```go
import "github.com/AImager/gocache/config"
import "github.com/go-redis/redis/v8"

func main() {
    a, b := 1, 3
	c := &Cache{client: redis.NewClient(&redis.Options{
		Addr: ":6379",
    })}

    // decorate get method, auto set cache
	decoratedGetFunc := getDb
	c.CacheWithFunc(context.TODO(), config.CacheConfig{
		Key:    fmt.Sprintf("test_cache:a:%d:b:%d", a, b),
		Expire: 300,
	}, &decoratedGetFunc, getDb)
    decoratedGetFunc(context.TODO(), a, b)

    // decorate update method, auto del cache
	decoratedUpdateFunc := updateDb
	c.CacheDelWithFunc(context.TODO(), config.CacheDelConfig{
		Key: fmt.Sprintf("test_cache:a:%d:b:%d", a, b),
	}, &decoratedUpdateFunc, updateDb)
	decoratedUpdateFunc(context.TODO(), a, b)
}

func getDb(ctx context.Context, a int, b string) (c int, err error) {
	return 1, nil
}

func updateDb(ctx context.Context, a int, b int) (err error) {
	return nil
}
```

## Contributing

PRs accepted.

## License

MIT © AImager