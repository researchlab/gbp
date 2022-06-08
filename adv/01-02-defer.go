package main

import "fmt"

func main() {
	fmt.Println(DeferFunc1(1))
	fmt.Println(DeferFunc2(1))
	fmt.Println(DeferFunc3(1))
}

// t 是全局变量， 所以返回t 赋值给t ,此时t= i =1
// 然后执行defer 语句, t += 3, t = 4
// 最后返回 t = 4
func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

// return t 将t的值赋值给返回的临时变量，此时值为i = 1
// 在函数返回之前 执行defer t = 3 + 1 = 4, 但是此时t 的值不会在赋值给返回的临时变量
// 最后返回的还是之前的临时变量值 1
func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

// 因为t 是全局变量， 所以
// return 2 先将2 赋值给t
// 在函数返回之前  先执行defer  t += i
// i = 1 所以 t = 2+1 = 3
func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

// output
// 4
// 1
// 3
