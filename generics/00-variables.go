package main

import "fmt"

// 定义泛型变量
// type 泛型变量名称   泛型参数列表（约束条件） 泛型类型
type Slice[T int | float64 | string] []T

type Map[KEY int | string, VALUE string | float64] map[KEY]VALUE

type Struct[T string | int | float64] struct {
	Title   string
	Content T
}

func main() {

	// 泛型初始化时，要将T替换为实参类型
	var s1 Slice[int] = []int{1, 2, 3}
	s2 := Slice[int]{1, 2, 3}

	fmt.Println("s1=", s1, " s2=", s2)

	var m1 Map[int, string] = map[int]string{1: "hello"}
	m2 := Map[int, string]{1: "hello"}

	fmt.Println("m1=", m1, " m2=", m2)

	var st Struct[float64]
	st.Title = "hello"
	st.Content = 3.12

	st2 := Struct[string]{Title: "hello", Content: "shell"}

	fmt.Println("st=", st, " st2=", st2)
}
