package main

import (
	"log"
	"time"

	"github.com/google/gops/agent"
)

// https://github.com/google/gops
func main() {
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Hour)
}
