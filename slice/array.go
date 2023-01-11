package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3, 4, 5}
	s1 := array[:2:2]
	s2 := array[2:]
	fmt.Println(s1) //[1 2]
	fmt.Println(s2) //[3 4 5]
	s1 = append(s1, 999) // 因为 len(s1) > cap(s1) s1 切换新的底层数组
	fmt.Println(array) // [1 2 3 4 5])
	fmt.Println(s1)    // [1 2 999]
	fmt.Println(s2)    // [3 4 5]
}
