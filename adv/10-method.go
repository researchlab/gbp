package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "good" {
		talk = "you are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	// 注意是*Student实现了接口People 而不是Student
	// 否则会报当前struct没有实现interface的提示
	// var peo People = Student{}
	var peo People = &Student{}
	think := "good"
	fmt.Println(peo.Speak(think))
}

// 实现接口的对象 要写对
// Student 实现了接口， 则初始化时 即可以写成Student{} 也可以写成&Student{}
// *Student 实现了接口， 初始化时只能写成&Student{}
