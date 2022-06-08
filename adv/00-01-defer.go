package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	// 实际是  defer calc("1", 1, calc("10", 1, 2))

	a = 0
	defer calc("2", a, calc("20", a, b))
	// 实际是  defer calc("2", 0, calc("20", 0, 2))
	// 特别注意， 这里ID calc("20",a, b) 因为是函数 所以先于后面的b = 1 执行
	// 所以这个函数执行时 b 还是等于2
	b = 1
}

/*
defer执行顺序
defer 函数中的参数 是先赋值的

注意到defer执行顺序和值传递
index:1肯定是最后执行的，但是index:1的第三个参数是一个函数，
所以最先被调用calc("10",1,2)==>10,1,2,3
执行index:2时,与之前一样，需要先调用calc("20",0,2)==>20,0,2,2
执行到b=1时候开始调用，index:2==>calc("2",0,2)==>2,0,2,2
最后执行index:1==>calc("1",1,3)==>1,1,3,4

10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4
*/
