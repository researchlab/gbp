package main

import (
	"fmt"
	"time"
)

func main() {
	r := Racer("a", "b")
	fmt.Println("main:", r)
	time.Sleep(5 * time.Second)
}

func Racer(a, b string) string {
	// select 则允许你同时在 多个 channel 等待。第一个发送值的 channel「胜出」，case 中的代码会被执行。
	select {
	case <-ping(a, 0):
		return a
	case <-ping(b, 1):
		return b
	}
}

func ping(url string, n int) chan bool {
	ch := make(chan bool)
	go func() {
		fmt.Println("url:", url)
		//time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		time.Sleep(time.Duration(n) * time.Second)
		ch <- true
	}()
	return ch
}
