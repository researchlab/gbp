package main

import (
	"fmt"
	"reflect"
)

// 说明多个子panic 可以统一被顶部的defer recover捕获到
// 并且， recover 只会保留最后一个panic的error
// 最开始的panic 将被丢弃
func main() {
	Foo()
	fmt.Println("-----------")
	Bar()
}

func Foo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("00", err)
		} else {
			fmt.Println("01 fatal")
		}
	}()
	defer func() {
		panic("defer panic")
	}()
	panic("panic")
}

/*
output
00 defer panic
*/
func Bar() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("+++++")
			f := err.(func() string)
			// 这里的err 是一个函数地址
			// 通过reflect.TypeOf(err).Kind() 获得err的类型是func
			fmt.Println(err, f(), reflect.TypeOf(err).Kind().String())
		} else {
			fmt.Println("fatal")
		}
	}()
	defer func() {
		panic(func() string {
			return "defer panic"
		})
	}()
	panic("panic")
}

// ++++
// 0x108a460 defer panic func
