package adapter

import (
	"context"

	"github.com/AImager/gocache/config"
	"github.com/gomodule/redigo/redis"
)

type redigoAdapter struct {
	client redis.Conn
}

func (g *redigoAdapter) Do(ctx context.Context, command string, args ...interface{}) (reply interface{}, err error) {
	return g.client.Do(command, args...)
}

func GetRedigoClient(cc config.ClientConfig) (config.ClientI, error) {
	conn, err := redis.Dial("tcp", cc.Addr)
	client := redigoAdapter{client: conn}
	return &client, err
}
