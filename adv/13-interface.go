package main

import "fmt"

func main() {
	var x *int = nil
	Foo(x)
	// non-empty
	var y interface{} = nil
	Foo(y)
	// empty
}

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty")
		return
	}
	fmt.Println("none-empty")
	fmt.Printf("%+v\n", x) // <nil>
}

// interface 有两种， 一种是empty 什么都没有
// 一种是none-empty 里面会记录 方法集， 数据类型, 上述案例中就记录了*int这个数据类型
// 所以 x 是 none-empty 的interface
