package main

import "fmt"

func main() {
	i := a()
	fmt.Println(i)
}

func a() int {
	var i int
	defer func() {
		i++
	}()
	return i
}
