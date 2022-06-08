package main

import "fmt"

// 当defer被声明时，其参数就会被实时解析

func main() {
	Foo()
}
func Foo() {
	i := 1
	defer fmt.Println(i)
	i++
	defer fmt.Println(i)
	i = 18
	defer fmt.Println(i)
	return
}

// output
// 18 2 1
