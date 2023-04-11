package test

import (
	"fmt"
	"testing"
)

func Test02Array(t *testing.T) {
	// t.Fatal("not implemented")
	arr := [3]int{0, 1, 2}
	f := func(v [3]int) {
		v[0] = 100
	}
	f(arr)           // no modify to arr
	fmt.Println(arr) // [0 1 2]
}
