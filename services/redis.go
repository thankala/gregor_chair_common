package services

import (
	"github.com/redis/go-redis/v9"
	"github.com/thankala/gregor_chair_common/configuration"
	"golang.org/x/net/context"
	"time"
)

type RedisStore struct {
	client  *redis.Client
	context *context.Context
}

func NewRedisStore(opts ...configuration.RedisOptFunc) *RedisStore {
	ctx := context.Background()
	options := configuration.DefaultRedisOpts()
	for _, opt := range opts {
		opt(options)
	}
	return &RedisStore{
		client: redis.NewClient(&redis.Options{
			Addr:     options.Addr,
			Password: options.Password,
			DB:       options.Db,
		}),
		context: &ctx,
	}
}

func (r *RedisStore) Ping() error {
	status := r.client.Ping(*r.context)
	return status.Err()
}

func (r *RedisStore) Load(key string) ([]byte, error) {
	val, err := r.client.Get(*r.context, key).Result()
	return []byte(val), err
}

func (r *RedisStore) Store(key string, value []byte) error {
	err := r.client.Set(*r.context, key, value, 0).Err()
	return err
}

func (r *RedisStore) Remove(key string) error {
	_, err := r.client.Del(*r.context, key).Result()
	return err
}

func (r *RedisStore) AcquireLock(key string, value string) (bool, error) {
	res, err := r.client.SetNX(*r.context, key, value, 1*time.Second).Result()
	if err != nil {
		return false, err
	}
	return res, nil
}

func (r *RedisStore) ReleaseLock(key string) error {
	_, err := r.client.Del(*r.context, key).Result()
	return err
}
