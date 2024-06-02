package main

import (
	"context"
	rdb2 "go-auth/pkg/rdb"
	"time"
)

func main() {
	rdb := rdb2.NewRdbCli(&rdb2.RedisOption{
		Password: "12345678",
	})

	result, err := rdb.Set(context.Background(), "a", 1, time.Hour).Result()
	if err == nil {
		println(result)
	}

	get := rdb.Get(context.Background(), "bbb")
	if get.Err() == nil {
		println(get.Val())
	} else {
		println(get.Err().Error())
	}
}
