package main

import "fmt"

func main() {
	c := make(chan int)
	go func(c chan int) {
		close(c)
		fmt.Println("goroutine exit")
	}(c)
	data, ok := <-c
	if !ok {
		fmt.Println("received close channel, data is:", data)
	}
	// invoke panic by closed again
	close(c)
	fmt.Println("main exit")
}

// output
//goroutine exit
//received close channel, data is: 0
//panic: close of closed channel
