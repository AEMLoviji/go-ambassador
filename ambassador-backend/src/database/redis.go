package database

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	Cache        *redis.Client
	CacheChannel chan string
)

func SetupRedis() {
	Cache = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})
}

func SetupCacheChannel() {
	CacheChannel = make(chan string)

	go func() {
		for {
			key := <-CacheChannel

			Cache.Del(context.Background(), key)

			// as if releasing cache do hard work
			time.Sleep(1 * time.Second)

			fmt.Printf("Cache cleared %s\n", key)
		}
	}()
}

func ClearCache(keys ...string) {
	for _, key := range keys {
		CacheChannel <- key
	}
}
