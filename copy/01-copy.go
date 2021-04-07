package main

import "fmt"

func main() {
	//ShowSliceAddr(0, 0)
	//ShowSliceAddr(1, 0)
	//ShowSliceAddr(2, 0)
	//ShowSliceAddr(3, 0)
	//ShowSliceAddr(3, 3)
	// output
	//address of slice 0xc00000c080 ; the len is 0 and cap is 0
	//address of slice 0xc00000c0a0 ; the len is 1 and cap is 1
	//address of slice 0xc00000c0c0 ; the len is 2 and cap is 2
	//address of slice 0xc00000c0e0 ; the len is 3 and cap is 3
	//address of slice 0xc00000c100 ; the len is 3 and cap is 3

	//appendSlice()
	copySlice()
}

func ShowSliceAddr(l, c int) {
	var s []int
	if l > 0 {
		s = make([]int, l)
	}
	if c > 0 {
		s = make([]int, l, c)
	}
	fmt.Printf("address of slice %p ; the len is %d and cap is %d \n", &s, len(s), cap(s))
}

func appendSlice() {
	var s []int
	fmt.Printf("address of slice %p ; the len is %d and cap is %d \n", &s, len(s), cap(s))
	s = append(s, 1)
	fmt.Printf("address of slice %p ; the len is %d and cap is %d \n", &s, len(s), cap(s))
	_s := append(s, 1)
	fmt.Printf("address of slice %p ; the len is %d and cap is %d \n", &_s, len(_s), cap(_s))

	//address of slice 0xc00000c080 ; the len is 0 and cap is 0
	//address of slice 0xc00000c080 ; the len is 1 and cap is 1
	//address of slice 0xc00000c0a0 ; the len is 2 and cap is 2
}

func copySlice() {
	var s, ss []int
	fmt.Printf("address of slice %p ; the len is %d and cap is %d \n", &s, len(s), cap(s))
	fmt.Printf("address of slice %p ; the len is %d and cap is %d \n", &ss, len(ss), cap(ss))
	//address of slice 0xc00000c080 ; the len is 0 and cap is 0
	//address of slice 0xc00000c0a0 ; the len is 0 and cap is 0

	// func copy(dst, src []Type) int
	copy(s, ss)
	fmt.Printf("address of slice %p ; the len is %d and cap is %d \n", &s, len(s), cap(s))
	fmt.Printf("address of slice %p ; the len is %d and cap is %d \n", &ss, len(ss), cap(ss))
	//address of slice 0xc00000c080 ; the len is 0 and cap is 0
	//address of slice 0xc00000c0a0 ; the len is 0 and cap is 0

	s = append(s, 1)
	copy(ss, s)
	fmt.Printf("address of slice %p ; the len is %d and cap is %d  and the slice is %v\n", &s, len(s), cap(s), s)
	fmt.Printf("address of slice %p ; the len is %d and cap is %d  and the slice is %v\n", &ss, len(ss), cap(ss), ss)
	//address of slice 0xc00000c080 ; the len is 1 and cap is 1 and the slice is [1]
	//address of slice 0xc00000c0a0 ; the len is 0 and cap is 0 and the slice is []

	ss = make([]int, len(s))
	copy(ss, s)
	fmt.Printf("address of slice %p ; the len is %d and cap is %d  and the slice is %v\n", &s, len(s), cap(s), s)
	fmt.Printf("address of slice %p ; the len is %d and cap is %d  and the slice is %v\n", &ss, len(ss), cap(ss), ss)
	//address of slice 0xc00000c080 ; the len is 1 and cap is 1 and the slice is [1]
	//address of slice 0xc00000c0a0 ; the len is 0 and cap is 0 and the slice is [1]
}

func copyMap() {
	var m, mm map[string]string
	fmt.Printf("address of map %p ; the len is %d\n", &m, len(m))
	fmt.Printf("address of map %p ; the len is %d\n", &mm, len(mm))

	m = make(map[string]string)
	mm = make(map[string]string)
	fmt.Printf("address of map %p ; the len is %d\n", &m, len(m))
	fmt.Printf("address of map %p ; the len is %d\n", &mm, len(mm))

	// copy(m, mm)
	// build error
	// arguments to copy must be slices; have map[string]string, map[string]string
}
