package main

import (
	"fmt"
	"reflect"
)

func main() {
	sn1 := struct {
		age  int
		name string
	}{age: 10, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 10, name: "qq"}
	sn3 := struct {
		name string
		age  int
	}{age: 10, name: "qq"}
	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}
	// invalid operation: sn1 == sn3 (mismatched types
	// 结构体比较，结构体字段要相同， 且顺序要相同，否则就是类型不对;
	// if sn1 == sn3 {
	// 	fmt.Println("sn1 == sn3")
	// }
	_ = sn3
	sm1 := struct {
		age int
		m   map[string]string
		i   interface{}
	}{age: 11, m: map[string]string{"a": "1"}, i: "interface"}
	sm2 := struct {
		age int
		m   map[string]string
		i   interface{}
	}{age: 11, m: map[string]string{"a": "1"}, i: "interface"}
	// invalid operation: sm1 == sm2 (struct containing map[string]string cannot be compared)
	// if sm1 == sm2 {
	// 	fmt.Println("sm1 == sm2")
	// }
	// 结构体中包含map, slice 是不能直接比较的,可以通过反射比较
	// 结构体中如果是interface , 且interface也是简单类型， 也可以直接比较
	if reflect.DeepEqual(sm1, sm2) {
		fmt.Println("sm1 == sm2")
	}
}

// 结构体比较
// 结构体字段要相同
// 结构体字段顺序要相同
// 结构体中如果只是简单类型，或者包含interface类型，但是Interface类型是简单值 可以直接比较
// 结构体如果包含map slice 等类型则需要用反射reflect.DeepEqual来比较
