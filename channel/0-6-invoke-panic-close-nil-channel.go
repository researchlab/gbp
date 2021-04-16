package main

func main() {
	var c chan int = nil
	close(c)
}

//output
// panic: close of nil channel
