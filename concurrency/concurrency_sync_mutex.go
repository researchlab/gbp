/**
*     File: con01.go
*	 Brief: 并发执行的数据同步问题
*	 Descr: 通过加锁Mutex.Clock解决并发过程中的数据同步问题
*
*   Author: Hong Li, leehongitrd@163.com
*   Github: https://github.com/researchlab
*  Created: 2016-02-18 01时36分37秒
* Modified: 2016-02-18 01时36分37秒
**/

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var total_tickets int32 = 10
var mutex = &sync.Mutex{}

func sell_tickets(wg *sync.WaitGroup, i int) {

	for total_tickets > 0 {

		mutex.Lock()
		// 如果有票就卖
		if total_tickets > 0 {
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
			// 卖一张票
			total_tickets--
			fmt.Println("id:", i, " ticket:", total_tickets)
		}
		mutex.Unlock()
	}
	wg.Done()
}

func main() {

	// 设置真正意义上的并发
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := sync.WaitGroup{}

	wg.Add(5)
	// 生成随机种子
	rand.Seed(time.Now().Unix())

	// 并发5个goroutine来卖票
	for i := 0; i < 5; i++ {
		go sell_tickets(&wg, i)
	}

	wg.Wait()
	// 退出时打印还有多少票
	fmt.Println(total_tickets, "done")
}
