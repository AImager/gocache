package cache

import (
	"errors"

	"github.com/AImager/gocache/adapter"
	"github.com/AImager/gocache/config"
)

func GetClient(c config.ClientConfig) (config.ClientI, error) {
	switch c.ClientType {
	case config.Goredis:
		return adapter.GetGoredisClient(c)
	case config.Redigo:
		return adapter.GetRedigoClient(c)
	default:
		return nil, errors.New("unsupport client type")
	}
}
