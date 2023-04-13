package main

import "fmt"

// 01 类型嵌套
type MyStruct[S int | string, P map[S]string] struct {
	Name    string
	Content S
	Job     P
}

// 02 泛型变量嵌套
type QSlice[T int | string] []T

// 注意
// 这里P的取值范围必须在 上面QSlice T之内，否则报错
// P does not satisfy int | string
type QStruct[P int | string, V QSlice[P]] struct {
	Name  P
	Title V
}

// 02.1
type Slice1[T int | float64 | string] []T

// 注意: Slice2[T int|string] 中T的类型约束必须在Slice1 的取值范围里，否则报错
// 如 Slice2[T int|bool] Slice1[T] 就不行， 因为bool 不在Slice1的T int|float64|string中
type Slice2[T int | string] Slice1[T]

// Slice3 继承自Slice2 但是 float64 在Slice1, 尽管Slice2 继承自Slice1
// 但是还是会报错， 因为泛型继承时单一的继承， Slice3不能通过Slice2 间接继承Slice1
// type Slice3[T float64 | int] Slice2[T]

func main() {

	var ms = MyStruct[int, map[int]string]{
		Name:    "shell",
		Content: 1,
		Job:     map[int]string{1: "s"},
	}
	ms2 := MyStruct[string, map[string]string]{
		Name: "Shell", Content: "hello", Job: map[string]string{"aa": "ss"},
	}
	fmt.Printf("ms=%+v, ms2=%+v\n", ms, ms2)

	qs := QSlice[int]{1}

	qms := QStruct[string, QSlice[string]]{Name: "string", Title: []string{"stite"}}

	fmt.Printf("qs=%+v, qms=%+v\n", qs, qms)

	// 02.1
	ms021 := Slice1[int]{1, 2}
	// Slice2其实就是继承和实现了Slice1，
	// 也就是说Slice2的类型参数约束的取值范围，必须是在Slice1的取值范围里。
	ms022 := Slice2[string]{"hello"}
	fmt.Printf("ms021=%+v, ms022=%+v\n", ms021, ms022)
}
