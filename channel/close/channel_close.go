package main

func main() {
	in := make(chan int, 2)
	in <- 1
	close(in)
	for _ = range in {
	}

	for {
		break
	}
}
