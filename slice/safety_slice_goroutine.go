package main

import (
	"fmt"
	"sync"
)

var w sync.WaitGroup
var m sync.Mutex

func sliceSafety() {
	var s []int
	var sum int
	fmt.Printf("----------: len(s): %d, cap(s): %d, s: %v \n", len(s), cap(s), s)
	for i := 0; i < 10; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			m.Lock()
			defer m.Unlock()
			sum++
			s = append(s, i)
			fmt.Printf("==========i: %d: len(s): %d, cap(s): %d, s: %v \n", i, len(s), cap(s), s)
		}(i)
	}
	w.Wait()
	fmt.Println(sum)
	fmt.Println(s, len(s))
}

func main() {
	sliceSafety()
}

/**
从结果可以看到，在加了锁之后，切片s中相同索引下存放的值总是相同的，没有遭到破坏，即加锁解决了线程安全的问题。至于是用互斥锁sync.Mutex还是用读写锁sync.RWMutex，这个看具体情况而定，如果读的场景远大于写的场景，用读写锁性能更好，因为读写锁又叫读写分离锁，在并发读的情况下不加锁，只有在并发写的情况下才加锁。

*/
