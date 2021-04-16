package main

func main() {
	var c chan int = nil
	c <- 100
}

// fatal error: all goroutines are asleep - deadlock!
