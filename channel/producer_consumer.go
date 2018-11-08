package channel

const (
	N = 100000
)

func Producer(out chan<- int) {
	for i := 1; i < N; i++ {
		out <- i
	}
	close(out)
}

func Consumer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x
	}
	close(out)
}
