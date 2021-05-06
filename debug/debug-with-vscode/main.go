package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	client, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	result, err := client.Do("SET", "key1", "value1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", result)

	result, err = client.Do("GET", "key1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", result)
}
