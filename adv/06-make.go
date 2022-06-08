package main

import "fmt"

func main() {
	s := make([]int, 3)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	// [0 0 0 1 2 3]

	// right
	ss := make([]int, 0)
	ss = append(ss, 1, 2, 3)
	fmt.Println(ss)
	// [1 2 3]
}
