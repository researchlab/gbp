package main

import (
	"fmt"

	"github.com/researchlab/gbp/debug/debug-local/utils"
)

func main() {
	fmt.Println("restart")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(utils.Add(100, 2000))
}
