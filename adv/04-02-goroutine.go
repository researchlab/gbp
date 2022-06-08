package main

import (
	"fmt"
	"sync"
)

func main() {
	wg1()
	// wg2()
	// wg3()
}

/*
panic: sync: negative WaitGroup counter
传递的wg 是一个副本
*/
func wg1() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int, wg sync.WaitGroup) {
			fmt.Println(i)
			wg.Done()
		}(i, wg)
	}
	wg.Wait()
}

// 不传wg  是可以的
func wg2() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// 要传wg 参数 就要传递指针
func wg3() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			fmt.Println(i)
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
}
