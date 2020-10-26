package redis

import (
	redisLib "github.com/go-redis/redis"
	"github.com/grokify/gostor"
)

type Client struct {
	redisClient *redisLib.Client
}

func NewClient(cfg gostor.Config) *Client {
	return &Client{
		redisClient: redisLib.NewClient(NewRedisOptions(cfg))}
}

func (client Client) SetString(key, val string) error {
	// For context, see https://github.com/go-redis/redis/issues/582
	// ctx, _ := context.WithTimeout(context.TODO(), time.Second)
	return client.redisClient.Set(key, val, 0).Err()
}

func (client Client) GetString(key string) (string, error) {
	// ctx, _ := context.WithTimeout(context.TODO(), time.Second)
	return client.redisClient.Get(key).Result()
}

func (client Client) GetOrEmptyString(key string) string {
	// ctx, _ := context.WithTimeout(context.TODO(), time.Second)
	if val, err := client.redisClient.Get(key).Result(); err != nil {
		return ""
	} else {
		return val
	}
}

func NewRedisOptions(cfg gostor.Config) *redisLib.Options {
	return &redisLib.Options{
		Addr:     cfg.HostPort(),
		Password: cfg.Password,
		DB:       cfg.CustomIndex}
}
