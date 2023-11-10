package rediser

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
)

func subscribe(ctx context.Context, rdb *redis.Client) {
	sub := rdb.Subscribe(ctx, "testchann")
	defer sub.Close()
	log.Println("strat sub")
	for {
		mes, err := sub.ReceiveMessage(ctx)
		if err != nil {
			return
		}
		log.Println("mes=:", mes)
	}
}

func pubscribe(ctx context.Context, rdb *redis.Client) {
	for i := 0; i < 10; i++ {
		rdb.Publish(ctx, "testchann", "test"+strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
