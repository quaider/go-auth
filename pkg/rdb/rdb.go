package rdb

import (
	"github.com/redis/go-redis/v9"
	"time"
)

var rdbCli redis.UniversalClient

func Client() redis.UniversalClient {
	return rdbCli
}

func NewRdbCli(option *RedisOption) redis.UniversalClient {
	if rdbCli != nil {
		return rdbCli
	}

	minIdleConns := option.MinIdleConns
	if minIdleConns == 0 {
		minIdleConns = 1
	}

	maxIdleConns := option.MaxIdleConns
	if maxIdleConns == 0 {
		maxIdleConns = 8
	}

	opts := &redis.UniversalOptions{
		MasterName:   option.MasterName,
		Addrs:        option.Addrs,
		DB:           option.Database,
		Password:     option.Password,
		PoolSize:     option.PoolSize,
		MaxRetries:   option.MaxRetries,
		ReadTimeout:  time.Duration(option.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(option.WriteTimeout) * time.Second,
		DialTimeout:  time.Duration(option.DialTimeout) * time.Second,
		PoolTimeout:  time.Duration(option.PoolTimeout) * time.Second,
		MinIdleConns: minIdleConns,
		MaxIdleConns: maxIdleConns,
	}

	rdbCli = redis.NewUniversalClient(opts)

	return rdbCli
}
