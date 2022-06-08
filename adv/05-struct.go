package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
}

/*
考点：go的组合继承

解答：这是Golang的组合模式，可以实现OOP的继承。
被组合的类型People所包含的方法虽然升级成了外部类型Teacher
这个组合类型的方法（一定要是匿名字段），
但它们的方法(ShowA())调用时接受者并没有发生变化。
此时People类型并不知道自己会被什么类型组合，
当然也就无法调用方法时去使用未知的组合者Teacher类型的功能。

showA
showB
*/
