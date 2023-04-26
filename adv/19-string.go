package main

// * string可以使用 `:` 来做字符串截取
//  **注意**：这里和[切片slice](../lesson13)的截取有区别
//    * 字符串截取后赋值给新变量，对新变量的修改不会影响原字符串的值
//   * 切片截取后复制给新变量，对新变量的修改会影响原切片的值

import "fmt"

func main() {
	s := "abc"
	fmt.Println(len(s)) // 3
	s1 := s[:]
	s2 := s[:1]
	s3 := s[0:]
	s4 := s[0:2]
	s5 := s[0:2]
	// update s5 and s  不影响s1,2,3,4
	s5 = "sss"
	s = "def"
	fmt.Println(s, s1, s2, s3, s4, s5) // def abc a abc ab sss
}
