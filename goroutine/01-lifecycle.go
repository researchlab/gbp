package main

import (
	"log"
	"time"
)

// 结论: 当父协程是main协程时，父协程退出，父协程下的所有子协程也会跟着退出；当父协程不是main协程时，父协程退出，父协程下的所有子协程并不会跟着退出（子协程直到自己的所有逻辑执行完或者是main协程结束才结束）
func main() {
	defer trace("main")()
	go goroutine()
	time.Sleep(20 * time.Second)
}

func goroutine() {
	defer trace("goroutine")()
	go goroutinue1()
	go goroutinue2()

}

func goroutinue1() {
	for i := 0; i < 3; i++ {
		log.Printf("this is goroutine1-%d", i)
		time.Sleep(2 * time.Second)
	}
}

func goroutinue2() {
	for i := 0; i < 3; i++ {
		log.Printf("this is goroutine2-%d", i)
		time.Sleep(2 * time.Second)
	}
}

//这个函数可以忽略，其作用只是辅助查看何时进入和退出某个方法
func trace(msg string) func() {
	log.Printf("enter %s", msg)
	start := time.Now()
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}
