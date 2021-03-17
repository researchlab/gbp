package main

import (
	"fmt"
	"time"
)

func main() {
	//	run()
	// defer 是针对 tfunc 返回的函数
	defer tfunc("main10")()
	go demo()
	time.Sleep(3 * time.Second)
	fmt.Println("main20")
}

func run() {
	defer trace("main00")
	fmt.Println("main01")
	fmt.Println("main02")
	demo()
	fmt.Println("main03")
}
func trace(msg string) {
	fmt.Println("enter ", msg)
}

func demo() {
	defer trace("demo00")
	fmt.Println("demo01")
}

func tfunc(msg string) func() {
	fmt.Println("start ", msg)
	return func() { fmt.Println("tfunc end") }
}
