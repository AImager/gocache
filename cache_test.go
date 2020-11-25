package cache

import (
	"context"
	"fmt"
	"testing"

	"github.com/AImager/gocache/config"
)

const (
	cacheKey = "test_cache:a:%d:b:%s"
)

func getRcConn() config.ClientI {
	client, _ := GetClient(config.ClientConfig{
		Addr:       "127.0.0.1:6379",
		ClientType: config.Goredis,
	})
	return client
}

func testDb(ctx context.Context, a int, b string) (c int, err error) {
	return 1, nil
}

func testUpdateDb(ctx context.Context, a int, b string) (err error) {
	return nil
}

func TestDb(t *testing.T) {
	a := 1
	b := "string123"

	c := &Cache{Client: getRcConn()}
	decoratedFunc := testDb
	c.CacheWithFunc(context.TODO(), config.CacheConfig{
		Key:    fmt.Sprintf(cacheKey, a, b),
		Expire: 220,
	}, &decoratedFunc, testDb)
	fmt.Println(decoratedFunc(context.TODO(), a, b))
}

func TestUpdateDb(t *testing.T) {
	a := 1
	b := "string123"

	c := &Cache{Client: getRcConn()}
	decoratedFunc := testUpdateDb
	c.CacheDelWithFunc(context.TODO(), config.CacheDelConfig{
		Key: fmt.Sprintf(cacheKey, a, b),
	}, &decoratedFunc, testUpdateDb)
	fmt.Println(decoratedFunc(context.TODO(), a, b))
}

type hand struct{}

func (h hand) testDb(ctx context.Context, a int, b string) (c int, err error) {
	return 1, nil
}

func (h hand) testUpdateDb(ctx context.Context, a int, b string) (err error) {
	return nil
}

func TestMethodDb(t *testing.T) {
	a := 1
	b := "string123"

	h := hand{}
	c := &Cache{Client: getRcConn()}
	decoratedFunc := h.testDb
	c.CacheWithFunc(context.TODO(), config.CacheConfig{
		Key:    fmt.Sprintf(cacheKey, a, b),
		Expire: 220,
	}, &decoratedFunc, testDb)
	fmt.Println(decoratedFunc(context.TODO(), a, b))
	fmt.Println(h.testDb(context.TODO(), a, b))
}

func TestMethodUpdateDb(t *testing.T) {
	a := 1
	b := "string123"

	h := hand{}
	c := &Cache{Client: getRcConn()}
	decoratedFunc := h.testUpdateDb
	c.CacheDelWithFunc(context.TODO(), config.CacheDelConfig{
		Key: fmt.Sprintf(cacheKey, a, b),
	}, &decoratedFunc, testUpdateDb)
	fmt.Println(decoratedFunc(context.TODO(), a, b))
	fmt.Println(h.testUpdateDb(context.TODO(), a, b))
}
