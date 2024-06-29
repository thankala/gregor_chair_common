package configuration

import (
	"strconv"
)

var (
	defaultRedisAddr     = "localhost:6379"
	defaultRedisPassword = "" // no password set
	defaultRedisDb       = 0  // use default DB
)

type RedisOpts struct {
	Addr     string
	Password string
	Db       int
}

func DefaultRedisOpts() *RedisOpts {
	return &RedisOpts{
		Addr:     defaultRedisAddr,
		Password: defaultRedisPassword,
		Db:       defaultRedisDb,
	}
}

type RedisOptFunc func(opts *RedisOpts)

func WithStoreAddr(storeAddr string) RedisOptFunc {
	if storeAddr == "" {
		return func(opts *RedisOpts) {}
	} else {
		return func(opts *RedisOpts) {
			opts.Addr = storeAddr
		}
	}
}

func WithStorePassword(storePassword string) RedisOptFunc {
	if storePassword == "" {
		return func(opts *RedisOpts) {}
	} else {
		return func(opts *RedisOpts) {
			opts.Password = storePassword
		}
	}
}

func WithStoreDb(storeDb string) RedisOptFunc {
	if storeDb == "" {
		return func(opts *RedisOpts) {}
	} else {
		return func(opts *RedisOpts) {
			db, err := strconv.Atoi(storeDb)
			if err != nil {
				return
			}
			opts.Db = db
		}
	}
}
