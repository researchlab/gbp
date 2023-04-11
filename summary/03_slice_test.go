package test

import (
	"fmt"
	"testing"
)

func Test03Slice(t *testing.T) {
	arr := []int{0, 1, 2}
	f := func(v []int) {
		v[0] = 100       // can modify origin array
		v = append(v, 4) // new memory allocated
		v[0] = 50        // no modify to origin array
	}
	f(arr)
	fmt.Println(arr)
}
