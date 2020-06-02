/**
*     File:		concurrency_sync_mutex.go
*	 Brief: 并发执行的数据同步问题
*	 Descr: 通过加锁Mutex.Clock解决并发过程中的数据同步问题
*
*   Author: Hong Li, leehongitrd@163.com
*   Github: https://github.com/researchlab
*  Created: 2016-02-18 01时36分37秒
* Modified:	2016-02-18 16时07分05秒
**/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	// 生成随机种子
	rand.Seed(time.Now().Unix())

	// 并发5个goroutine来卖票
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go SellTickets(&wg, i)
	}

	wg.Wait()
	// 退出时打印还有多少票
	fmt.Println(TotalTickets(), "done")
}

// total tickets
var total_tickets int32 = 10

// global mutex
var mutex = &sync.Mutex{}

func SellTickets(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	for total_tickets > 0 {
		sell(i)
	}
}

func sell(i int) {
	mutex.Lock()
	defer mutex.Unlock()

	// 如果有票就卖
	if total_tickets > 0 {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		// 卖一张票
		total_tickets--
		fmt.Println("id:", i, " ticket:", total_tickets)
	}
}

func TotalTickets() int32 {
	return total_tickets
}
