package test

import (
	"fmt"
	"testing"
)

func TestAnySlice(t *testing.T) {
	appendAnyF := func(t []any, toAppend ...any) []any {
		ret := append(t, toAppend...)
		return ret
	}

	emptySlice := []any{}
	slice2 := []any{"hello", "world"}

	// bug append slice as a element
	emptySlice = appendAnyF(emptySlice, slice2)
	fmt.Println(len(emptySlice), emptySlice) // only 1 element [[hello world]]

	emptySlice = []any{}
	emptySlice = appendAnyF(emptySlice, slice2...)
	fmt.Println(len(emptySlice), emptySlice) // [hello world]
}
