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

package lib

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var total_tickets int32 = 10
var mutex = &sync.Mutex{}

func SellTickets(wg *sync.WaitGroup, i int) {

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

func TotalTickets() int32 {
	return total_tickets
}
