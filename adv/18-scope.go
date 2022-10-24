package main

import "fmt"

func main() {
	err := Foo()
	fmt.Println("err:", err)
}

func Foo() (err error) {
	if err := Bar(); err != nil {
		//  这里如果直接返回 return 
		// 则会报错: ./18-scope.go:12:3: err is shadowed during return
		// 意思是return 时 全局的err变量被隐藏/屏蔽了, 因为局部有一个变量名也叫err
		// 必须使用 return err
		return err
	}
	return
}

func Bar() (err error) {
	return fmt.Errorf("Bar Error")
}
