package main

import "C"

// 代码通过import "C"语句启用CGO特性，主函数只是通过Go内置的println函数输出字符串，其中并没有任何和CGO相关的代码。虽然没有调用CGO的相关函数，但是go build命令会在编译和链接阶段启动gcc编译器，这已经是一个完整的CGO程序了
func main() {
	println("hello cgo")
}
