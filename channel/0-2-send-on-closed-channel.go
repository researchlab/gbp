package main

func main() {
	in := make(chan bool)
	close(in)
	go func() { <-in }()
	in <- true // panic: send on closed channel
}
