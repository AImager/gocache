package adapter

import (
	"context"

	"github.com/AImager/gocache/config"
	"github.com/go-redis/redis/v8"
)

type goredisAdapter struct {
	client *redis.Client
}

func (g *goredisAdapter) Do(ctx context.Context, command string, args ...interface{}) (reply interface{}, err error) {
	argsTemp := []interface{}{command}
	for _, arg := range args {
		argsTemp = append(argsTemp, arg)
	}
	cmd := g.client.Do(ctx, argsTemp...)
	return cmd.Val(), cmd.Err()
}

func GetGoredisClient(cc config.ClientConfig) (config.ClientI, error) {
	c := redis.NewClient(&redis.Options{
		Addr: cc.Addr,
	})
	return &goredisAdapter{client: c}, nil
}
