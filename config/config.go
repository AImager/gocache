package config

import "context"

type ClientI interface {
	Do(ctx context.Context, command string, args ...interface{}) (reply interface{}, err error)
}

type ClientConfig struct {
	CacheType  cacheType
	ClientType clientType
	Addr       string
}

type cacheType string

const (
	Redis cacheType = "redis"
)

type clientType string

const (
	Goredis clientType = "goredis"
	Redigo  clientType = "redigo"
)
