// test.go
package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	for i := 1; i <= 10; i++ {
		go func(gid int) {
			n := 0
			for {
				fmt.Println(time.Now().Format("2006-01-02 15:04:05"), gid, n)
				time.Sleep(time.Second)
			}
		}(i)
	}

	go func() {
		arr := 0
		p := uintptr(unsafe.Pointer(&arr))
		myfun1(p)
	}()

	for true {
		time.Sleep(time.Second)
	}
}

func myfun1(p uintptr) {
	arr := (*int)(unsafe.Pointer(p))
	*arr = 1
	fmt.Println(*arr)
	go myfun2()
	fmt.Println(*arr)
}

func myfun2() {
	fmt.Println("myfun2")
	myfun3()
}

func myfun3() {
	var p uintptr = 0
	arr := (*int)(unsafe.Pointer(p))
	*arr = 1
	fmt.Println(*arr)
}
