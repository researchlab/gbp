package main

import (
	"fmt"
	"sync"
)

func FibIter(a, b, n int) int {
	if n == 0 {
		return b
	}

	return FibIter(a+b, a, n-1)
}

func Fib(n int) int {
	return FibIter(1, 0, n)
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			r := Fib((i + 1) * 100)
			fmt.Println(r)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
