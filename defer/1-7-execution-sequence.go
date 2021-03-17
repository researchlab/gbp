package main

import "fmt"

// Slice array
type Slice []int

// NewSlice new slice
func NewSlice() Slice {
	return make(Slice, 0)
}

// Add add num
func (s *Slice) Add(elem int) *Slice {

	*s = append(*s, elem)
	fmt.Println(elem)
	return s
}

func main() {
	s := NewSlice()

	// defer 作用于最后一个函数
	defer s.Add(1).Add(2).Add(3)

	s.Add(4)

	fmt.Println(s)

	// output
	//  1
	//  2
	//  4
	//  [1 2 4]
	//  3
}
