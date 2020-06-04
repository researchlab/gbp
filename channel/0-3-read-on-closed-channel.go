package main

import "log"

func main() {
	in := make(chan int, 2)
	close(in)
	_, ok := <-in
	log.Printf("ok:%v", ok) // ok:false
}
