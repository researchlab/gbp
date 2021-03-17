package main

import "fmt"

func main() {
	demo()
}

func demo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover:", r)
		}
	}()
	defer fmt.Println("demo")
	panic("panic")
	defer fmt.Println("finished")
}
