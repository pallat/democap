package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"time"

	redis "github.com/redis/go-redis/v9"
)

func main() {
	masterSet()
	masterGet()
	slave()
}

func masterSet() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	defer rdb.Close()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatal(err)
	}

	val := "555"
	if err := rdb.Set(context.Background(), "key1", val, 0).Err(); err != nil {
		log.Panicf("set %v not success: %s\n", val, err)
	}

	all := 1
	intCmd := rdb.Wait(context.Background(), all, 3*time.Second)
	if int(intCmd.Val()) != all {
		slog.Error("not guarantee consistency")
	}

	fmt.Printf("set %q to master successfully\n", val)
}

func masterGet() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	defer rdb.Close()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatal(err)
	}

	cmd := rdb.Get(context.Background(), "key1")
	fmt.Printf("latest value from master: %v\n", cmd.Val())
}

func slave() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6479", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	defer rdb.Close()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatal(err)
	}

	cmd := rdb.Get(context.Background(), "key1")
	fmt.Printf("latest value from slave: %v\n", cmd.Val())
}
