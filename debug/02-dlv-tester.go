package main

import (
	"fmt"
	"runtime"
	"time"
)

func deadloop() {
	for {
	}
}

func worker() {
	for {
		fmt.Println("worker is running")
		time.Sleep(time.Second * 1)
	}
}

func main() {
	fmt.Printf("There are %d cores.\n", runtime.NumCPU())

	go worker()

	go deadloop()

	i := 3
	for {
		fmt.Printf("main is running, i=%d\n", i)
		i--
		if i == 0 {
			runtime.GC()
		}

		time.Sleep(time.Second * 1)
	}
}
