package main

import "fmt"

func main() {
	// list := new([]int)
	list := make([]int, 0)
	list = append(list, 1)
	fmt.Println(list)
	arr := new([5]int)
	arr[0] = 10
	fmt.Println(arr)
}

// new 返回的是一个指针，不是一个切片, 所以切片初始化不能使用new
// 但是数组可以使用new初始化
