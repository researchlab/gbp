package main

import (
	"fmt"
	"sync"
)

// golang 两个协成交替打印1-100的奇数偶数

var POOL = 100

func goroutine1(wg *sync.WaitGroup, p chan int) {
	for i := 1; i <= POOL; i++ {
		p <- i
		if i%2 == 1 {
			wg.Done()
			fmt.Println("goroutine 1 ", i)
		}
	}
}

func goroutine2(wg *sync.WaitGroup, p chan int) {
	for i := 1; i <= POOL; i++ {
		<-p
		if i%2 == 0 {
			wg.Done()
			fmt.Println("goroutine 2 ", i)
		}
	}
}

func main() {
	msg := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(POOL)
	go goroutine1(wg, msg)
	go goroutine2(wg, msg)
	wg.Wait()
}
