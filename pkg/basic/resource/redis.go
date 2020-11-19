package resource

import (
	"errors"

	"github.com/alonegrowing/cafe/pkg/sea/redis"
)

var (
	RedisClientNotInit = errors.New("redis client not init")
)

var defaultRedis map[string]*redis.Redis

func NewRedis(configs []redis.Config) error {
	if defaultRedis == nil {
		defaultRedis = make(map[string]*redis.Redis)
	}
	for _, conf := range configs {
		client, err := redis.NewRedis(&conf)
		if err != nil || client == nil {
			continue
		}
		defaultRedis[conf.ServerName] = client
	}
	return nil
}

func GetRedis(service string) (*redis.Redis, error) {
	if client, ok := defaultRedis[service]; ok {
		return client, nil
	}
	return nil, RedisClientNotInit
}
