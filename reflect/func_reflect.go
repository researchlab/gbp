/**
*     File:		func_reflect.go
*	 Brief: 通过反射“动态”调用方法
*	 Descr:
*
*   Author: Hong Li, leehongitrd@163.com
*   Github: https://github.com/researchlab
*  Created: 2016-02-17 13时08分23秒
* Modified:	2016-02-18 16时10分06秒
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

func (u User) Hello(m User) (int, string) {
	fmt.Println("Hello", m.Name, ", I'm ", u.Name)
	return m.Age + u.Age, u.Name
}

func main() {
	u := User{1, "Mike", 11}
	GetInfo(u)
}

func GetInfo(u interface{}) {
	m := User{2, "Json", 12}

	v := reflect.ValueOf(u)

	if v.Kind() != reflect.Struct {
		fmt.Println("type invalid")
		return
	}

	mv := v.MethodByName("Hello") //获取对应的方法
	if !mv.IsValid() {            //判断方法是否存在
		fmt.Println("Func Hello not exist")
		return
	}

	args := []reflect.Value{reflect.ValueOf(m)} //初始化传入等参数，传入等类型只能是[]reflect.Value类型
	res := mv.Call(args)
	fmt.Println(res[0], res[1])

}
