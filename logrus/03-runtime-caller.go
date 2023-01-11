package main

import (
	"fmt"
	"runtime"
)

func main() {
	foo()
	fmt.Println("boo----")
	boo()
	fmt.Println("coo----")
	coo()
	fmt.Println("doo----")
	doo()
}


func foo() {
	pc, file, line, ok := runtime.Caller(1)
	fmt.Println("pc:", pc)
	fmt.Println("file:", file)
	fmt.Println("line:", line)
	fmt.Println("ok:", ok)
}

func boo() {
	pc, file, line, ok := runtime.Caller(1)
	fmt.Println("pc:", pc)
	fmt.Println("file:", file)
	fmt.Println("line:", line)
	fmt.Println("ok:", ok)
}

func coo() {
	log()
}
func doo() {
	log()
}
func log() {
	pc, file, line, ok := runtime.Caller(2)
	fmt.Println("pc:", pc)
	fmt.Println("file:", file)
	fmt.Println("line:", line)
	fmt.Println("ok:", ok)
}
