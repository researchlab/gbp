package main

import "fmt"

// 闭包引用相同变量, 存在修改的分险
func test(x int) (func(), func()) {
	return func() {
			fmt.Println(x)
			x += 18
		}, func() {
			fmt.Println(x)
		}
}

func main() {
	a, b := test(100)
	a()
	b()
}

// output
// 100
// 118
