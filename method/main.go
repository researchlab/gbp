package main

import (
	"github.com/researchlab/gbp/method/pkg"
	"fmt"
)
func main() {
	r := pkg.New("user").Query()
	fmt.Println(r)
	scope01 := pkg.Scope()
	scope02 := pkg.Scope2()
	fmt.Println(scope01)
	fmt.Println(scope02)
}
