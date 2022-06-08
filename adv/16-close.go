package main

import "fmt"

// 闭包延迟赋值

func test() []func() {
	var funcs []func()
	for i := 0; i < 2; i++ {
		funcs = append(funcs, func() {
			fmt.Println(&i, i) // 闭包延迟赋值
		})
	}

	/*
		for i := 0; i < 2; i++ {
			x := i
			funcs = append(funcs, func() {
				fmt.Println(&x, x)
			})
		}
		output
		0xc0000b0008 0
		0xc0000b0010 1
	*/
	return funcs
}

func main() {
	funcs := test()
	for _, f := range funcs {
		f()
	}
}

// output
// 0xc000096008 2
// 0xc000096008 2
