package test

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test05Pointer(t *testing.T) {
	slice := []int{0, 1, 2}
	ptr := unsafe.Pointer(&slice[0]) // get array element:0 pointer

	slice = append(slice, 3)          // allocate new memory
	ptr2 := unsafe.Pointer(&slice[0]) // 两次取的地址是不一样的， 因为取的时候是值变量，不是引用变量

	// output: ptr is 1374390738968, ptr2 is 1374390788144, ptr==ptr2 result is false
	fmt.Println(fmt.Sprintf("ptr is %d, ptr2 is %d, ptr==ptr2 result is %v", ptr, ptr2, ptr == ptr2))
}
