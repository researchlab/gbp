package main

import (
	"fmt"
	"sync"
	"time"
)

// key必须支持==和!=比较，才能用作map的key。

// 因此切片slice，函数类型function，集合map，不能用作map的key

// map不是并发安全的，并发读写要加锁

// 通过匿名结构体声明一个变量counter, 变量中包含了map 和 sync.RWMutex
var counter = struct {
	sync.RWMutex
	m map[string]int
	n map[string]Info
	x map[string]string
}{m: make(map[string]int), n: map[string]Info{}, x: make(map[string]string)}

type Info struct {
	Name string
}

func read() {
	// 读取数据的时候用读锁
	counter.RLock()
	defer counter.RUnlock()
	fmt.Println(counter.m["Tony"])
	fmt.Println(counter.n["Name"])
	fmt.Println(counter.x["timer"])
}

func write() {
	// 写数据的时候用写锁
	counter.Lock()
	defer counter.Unlock()
	counter.m["Tony"]++ // map 可以直接更新int
	counter.n["Name"] = Info{Name: fmt.Sprintf("mike:%v", time.Now().Unix())}
	// counter.n["Name"].Name = "update"  // map 不能直接更新struct
	counter.x["timer"] = fmt.Sprintf("%v", time.Now().Unix()) // map 可以直接更新string
}

func main() {
	write()
	read()
	write()
	read()
}

// safety map 是通过RWMutex 读锁和写锁来控制， 可以通匿名struct来简化代码
// map中可以直接对简单对象赋值 如int, string
// map 中不能直接更新或赋值给 struct
