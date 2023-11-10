package rediser

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/go-redis/redis/v9"
)

func TestSubscribe(test testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	log.Println("redis clietn")
	defer rdb.Close()
	ctx := context.Background()

	go subscribe(ctx, rdb)
	time.Sleep(time.Second)
	// rdb.PSubscribe()
	// pubscribe(ctx, rdb)
	time.Sleep(time.Second * 5)
}
