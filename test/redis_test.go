package test

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func TestSetValue(t *testing.T) {
	err := rdb.Set(ctx, "key", "value", time.Second*10).Err()
	if err != nil {
		t.Error(err)
	}
}

func TestGetValue(t *testing.T) {
	result, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
