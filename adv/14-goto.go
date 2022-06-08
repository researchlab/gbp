package main

import "fmt"

func main() {
	Foo()
}

func Foo() {
	for i := 0; i < 10; i++ {
		/*
			 写成这样 会报错
			 goto loop jumps into block starting
				if i == 2 {
				goto LABEL
				}
				fmt.Println(i)
		*/
		// 必须写成下面的 if else 否则不行
		if i == 2 {
			// 关键字，goto跳转到某个位置，且只能在当前函数内跳转
			goto LABEL
		} else {
			fmt.Println(i)
		}
	}
LABEL:
	fmt.Println("程序结束")
}
