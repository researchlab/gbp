package channel

import "testing"

func producerconsumer(in chan int, out chan int) {
	go Producer(in)
	go Consumer(in, out)

	for x := range out {
		_ = x
	}
}

func TestNoBufferChan(t *testing.T) {
	in, out := make(chan int), make(chan int)
	producerconsumer(in, out)
}

func TestBufferChan(t *testing.T) {
	bufLen := 100
	in, out := make(chan int, bufLen), make(chan int, bufLen)
	producerconsumer(in, out)
}

func BenchmarkNoBufferChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in, out := make(chan int), make(chan int)
		producerconsumer(in, out)
	}
}

func BenchmarkBufferChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bufLen := 100
		in, out := make(chan int, bufLen), make(chan int, bufLen)
		producerconsumer(in, out)
	}
}
