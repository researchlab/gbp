package main

import (
	"fmt"
	"sync"
	"time"
)

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	// ch := make(chan interface{}) 这里会阻塞
	ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()
		for elem, value := range set.s {
			ch <- elem
			fmt.Println("Iter:", elem, value)
		}
		close(ch) // 如果这里没有close
		// 在读取完所有值之后继续读取 会报 fatal error: all goroutines are asleep - deadlock!

		set.RUnlock()
	}()
	return ch
}

func main() {
	th := threadSafeSet{
		s: []interface{}{
			"1", "2",
		},
	}
	v := th.Iter()
	time.Sleep(time.Second)
	fmt.Printf("%s len=%v\n", "ch", len(v))
	fmt.Println(<-v)
	fmt.Println(<-v)
	fmt.Println(<-v)
	fmt.Println(<-v)
	fmt.Println(v)
}

/*
output
Iter: 0 1
Iter: 1 2
ch len=2
0
1
<nil>
<nil>
0xc000124120

chan 初始化必须要指定长度，否则会阻塞
len(chan) 得到chan 的长度
chan 没有数据后， 再次读取 <-chan  会返回类型空值， 如interface{} -> nil
chan 使用之后 如果忘记close 则在读取完之后继续读取会报 fatal error 死锁
*/
