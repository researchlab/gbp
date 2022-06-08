package main

import (
	"fmt"
	"sync"
)

/*
golang互斥锁的两种实现
1.用Mutex实现
2.使用chan实现
*/

// Mutex

var num int
var mtx sync.Mutex
var wg sync.WaitGroup

func add() {
	mtx.Lock()
	defer mtx.Unlock()
	defer wg.Done()
	num += 1
}

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println("num:", num)
}
