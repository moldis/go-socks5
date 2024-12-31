package main

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/things-go/go-socks5"
)

func main() {
	// Create a SOCKS5 server
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // use default DB
	})
	res := rdb.Ping(context.Background())
	if res.Err() != nil {
		panic(res.Err())
	}

	server := socks5.NewServer(
		socks5.WithLogger(socks5.NewLogger(log.New(os.Stdout, "socks5: ", log.LstdFlags))),
		socks5.WithRedisClient(rdb),
	)

	// Create SOCKS5 proxy on localhost port 8000
	log.Println("running on: ", os.Getenv("SOCKS_ADDRESS"))
	if err := server.ListenAndServe("tcp", os.Getenv("SOCKS_ADDRESS")); err != nil {
		panic(err)
	}
}
