package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A:", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B:", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

/*
B: 9
A: 10
A: 10
A: 10
A: 10
A: 10
A: 10
A: 10
A: 10
A: 10
A: 10
B: 0
B: 1
B: 2
B: 3
B: 4
B: 5
B: 6
B: 7
B: 8
*/
