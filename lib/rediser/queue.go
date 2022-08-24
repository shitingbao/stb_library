package rediser

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v9"
)

func setQueue(rdb *redis.Client, key string, values ...interface{}) {
	rdb.LPush(context.Background(), key, values...)
}

func queueWatch(rdb *redis.Client, timeout time.Duration, keys ...string) {
	for {
		res := rdb.BLPop(context.Background(), timeout, keys...)

		log.Println("android===>Val:", res.Val())

		queues := res.Val()
		if len(queues) == 0 {
			continue
		}
		log.Println("queues==>:", queues)
	}
}
