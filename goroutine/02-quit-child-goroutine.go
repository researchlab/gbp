package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//problem()
	solution1()
}

/*
1. main 函数退出，所有协程退出
2. 协程无父子关系，即在父协程开启新的协程，若父协程退出，不影响子协程

解决方式
通过context上下文来解决，当然也可以通过channel管道来解决，context解决方式如solution
*/
func problem() {
	fmt.Println("main 函数 开始...")
	go func() {
		fmt.Println("父 协程 开始...")
		go func() {
			for {
				fmt.Println("子 协程 执行中...")
				timer := time.NewTimer(time.Second * 2)
				<-timer.C
			}
		}()
		time.Sleep(time.Second * 5)
		fmt.Println("父 协程 退出...")
	}()
	time.Sleep(time.Second * 10)
	fmt.Println("main 函数 退出")
}

// 需要等到子协程的逻辑执行完成之后，第二次开始时才能收到父协程发出的退出信号
// 需要注意的时 golang 的协程不像线程和进程可以强制被中断的，需要等goroutine自己中断
func solution1() {
	fmt.Println("main 函数 开始...")
	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		fmt.Println("父 协程 开始...")
		go func(ctx context.Context) {
			for {
				for {
					select {
					case <-ctx.Done():
						fmt.Println("子 协程 接受停止信号...")
						return
					default:
						fmt.Println("子 协程 执行中...")
						timer := time.NewTimer(time.Second * 5)
						<-timer.C
						fmt.Println("子 协程 休眠结束")
					}
				}
			}
		}(ctx)
		time.Sleep(time.Second * 2)
		fmt.Println("父 协程 退出...")
	}()
	time.Sleep(time.Second * 10)
	fmt.Println("main 函数 退出")

}
