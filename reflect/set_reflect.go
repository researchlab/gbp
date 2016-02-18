/**
*     File:		set_reflect.go
*	 Brief: 通过反射修改结构体字段值
*	 Descr: 通过判断反射返回类型，元素是否能被修改，是否存在等成功修改相应等字段值
*
*   Author: Hong Li, leehongitrd@163.com
*   Github: https://github.com/researchlab
*  Created: 2016-02-17 13时06分33秒
* Modified:	2016-02-18 16时10分03秒
**/

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	u := User{1, "Mike", 11}
	SetInfo(&u)
	fmt.Println(u)
}

func SetInfo(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() { //判断是否为指针类型 元素是否可以修改
		fmt.Println("cannot set")
		return
	} else {
		v = v.Elem() //实际取得的对象
	}

	//判断字段是否存在
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("var invalid")
		return
	}

	//设置字段
	if f := v.FieldByName("Name"); f.Kind() == reflect.String {
		f.SetString("BY")
	}
}
