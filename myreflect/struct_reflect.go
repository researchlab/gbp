/**
*     File:		struct_reflect.go
*	 Brief: 简单使用反射提供接口对象信息
*	 Descr: 通过反射TypeOf/ValueOf/Field/NumMethod等方法获取接口对象的字段,类型和方法等信息
*
*   Author: Hong Li, leehongitrd@163.com
*   Github: https://github.com/researchlab
*  Created: 2016-01-13 11时04分41秒
* Modified:	2016-02-18 16时09分57秒
**/

package myreflect

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("Hello,Mike")
}

func StructReflectInit() {
	u := User{1, "Mike", 11}
	Info(u)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)         //获取接口的类型
	fmt.Println("Type:", t.Name()) //t.Name() 获取接口的名称

	if t.Kind() != reflect.Struct {
		fmt.Println("err: type invalid")
		return
	}

	v := reflect.ValueOf(o) //获取接口的值类型
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ { //NumField取出这个接口所有的字段数量
		f := t.Field(i)                                   //取得结构体的第i个字段
		val := v.Field(i).Interface()                     //取得字段的值
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val) //第i个字段第名称,类型,值
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}
