package main

import (
	"fmt"
	"sync"
)

var num int

func add(h chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	h <- 1
	num += 1
	<-h
}

func main() {
	ch := make(chan int, 1)
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add(ch, wg)
	}
	wg.Wait()
	fmt.Println("num:", num)
}

// 阻塞chan 只能用在goroutine 的异步任务中
// 如果直接用chan  一定要初始化一个缓冲channel, 否则阻塞死锁
