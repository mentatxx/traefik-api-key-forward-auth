package database

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func acquireLock(client *redis.Client) bool {
	lock, err := client.SetNX(context.Background(), redisLockKey, true, 10*time.Second).Result()
	if err != nil {
		fmt.Println("Error acquiring lock:", err)
		return false
	}
	return lock
}

func releaseLock(client *redis.Client) error {
	_, err := client.Del(context.Background(), redisLockKey).Result()
	return err
}
