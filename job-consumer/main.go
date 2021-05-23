package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

const key = "myJobQueue"

func main() {
	c := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	fmt.Println("Wating fo jobs on job queue", key)
	go func() {
		for {
			result, err := c.BLPop(c.Context(), 0*time.Second, key).Result()

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Excuting job", result[1])
		}
	}()

	select {}
}
