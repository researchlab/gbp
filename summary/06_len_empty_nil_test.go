package test

import (
	"fmt"
	"testing"
)

func Test06EmptyNil(t *testing.T) {
	var slice []int = nil
	fmt.Println(len(slice), cap(slice)) // 0 0

	var slice2 []int = []int{}
	fmt.Println(len(slice2), cap(slice2)) // 0 0

	var mp map[int]int = nil
	fmt.Println(len(mp)) // 0

	var mp2 map[int]int = map[int]int{}
	fmt.Println(len(mp2)) // 0
}
