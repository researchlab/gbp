package main

import (
	"errors"
	"fmt"
)

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(do bool) (err error) {
	if do {
		// 同名局部变量 替换了全局的变量， 所以这里的
		// err  虽然被赋值为ErrDidNotWork但是它是局部变量
		_, err := tryTheThing()
		if err != nil {
			err = ErrDidNotWork
		}
	}
	// 这里返回的err 是函数变量
	return err
}

func tryTheThing() (string, error) {
	return "", ErrDidNotWork
}

func main() {
	fmt.Println(DoTheThing(true))  // "<nil>" 因为同名的局部变量替换了全局变量
	fmt.Println(DoTheThing(false)) // "<nil>"
}
