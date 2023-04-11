package test

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestSizeofPtrBug(t *testing.T) {
	type CodeLocation struct {
		LineNo int64
		ColNo  int64
	}

	cl := &CodeLocation{10, 20}
	size := unsafe.Sizeof(cl) // 只会返回ptr 本身的大小，而非结构体自己的大小
	fmt.Println(size)         // always return 8 for point size
}

// 获得结构体指针的size, 可以通过反射来获取
func TestSizeofPtrWithoutBug(t *testing.T) {
	type CodeLocation struct {
		LineNo int64
		ColNo  int64
	}

	cl := &CodeLocation{10, 20}
	size := ValueSizeof(cl)
	fmt.Println(size) //16
}

func ValueSizeof(v any) uintptr {
	typ := reflect.TypeOf(v)
	if typ.Kind() == reflect.Pointer {
		return typ.Elem().Size()
	}
	return typ.Size()
}
