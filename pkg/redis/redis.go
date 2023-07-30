package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisPkg struct {
	client *redis.Client
}

func NewRedis(redisOpt *redis.Options) (*RedisPkg, error) {
	client := redis.NewClient(redisOpt)

	return &RedisPkg{
		client: client,
	}, nil
}

func (r *RedisPkg) SetJSONTTL(ctx context.Context, key string, value interface{}, TTL time.Duration) error {
	b, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("[SetJSON] error when Marshall, err: %+v", err)
	}

	return r.client.Set(ctx, key, string(b), TTL).Err()
}

func (r *RedisPkg) GetJSON(ctx context.Context, key string, data interface{}) error {
	res, err := r.client.Get(ctx, key).Result()

	if err != nil {
		return fmt.Errorf("[GetJSON] error when get data, err: %+v", err)
	}

	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return fmt.Errorf("[GetJSON] error when Unmarshal, err: %+v", err)
	}

	return nil
}
