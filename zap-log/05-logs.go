package main

import "github.com/researchlab/gbp/log/logs"

func main() {
	foo()
}

func foo() {
	logs.Error("logs error")
}
