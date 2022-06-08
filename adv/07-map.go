package main

import (
	"fmt"
	"sync"
	"time"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	// ua := new(UserAges)
	// ua.Add("mike", 10)
	// fmt.Println("mike age:", ua.Get("mike"))
	// output: panic: assignment to entry in nil map
	// new 并不会初始化struct里面的map

	// right := UserAges{
	// 	ages: make(map[string]int),
	// }
	// right.Add("mike", 10)
	// fmt.Println("mike age: ", right.Get("mike"))
	// // 单个读写是没问题的

	cur := UserAges{
		ages: make(map[string]int),
	}
	go func() {
		for i := 0; i < 1000; i++ {
			cur.Add(fmt.Sprintf("i%d", i), i)
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			cur.Get(fmt.Sprintf("i%d", i))
		}
	}()
	time.Sleep(time.Minute)
	// output fatal error: concurrent map read and map write

}

// new 一个struct, 并不会初始化struct 里面的map 等引用类型;
// 否则panic: assignment to entry in nil map
// struct 中使用map 起读写操作都要加锁，否则会出现fatal error

// 考点：map线程安全

// 解答：可能会出现fatal error: concurrent map read and map write. 修改一下看看效果
