package main

import "log"

func main() {
	in := make(chan int, 2)
	in <- 1
	close(in)
	for _ = range in {
	}

	for {
		break
	}

	_, ok := <-in
	log.Printf("ok:%v", ok)
}
