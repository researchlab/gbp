package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// golang 中结构体的深拷贝和浅拷贝

type User struct {
	Name string
}

// common
// DeepCopy copy
func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func main() {
	fmt.Println("引用类型, 直接赋值浅拷贝:")
	user1 := &User{Name: "why"}
	user2 := user1
	fmt.Printf("before user1:%v, user2:%v\n", user1, user2)
	user1.Name = "update"
	fmt.Printf("after user1:%v, user2:%v\n", user1, user2)

	fmt.Println("值类型, 直接赋值深拷贝:")
	user3 := User{
		Name: "why",
	}
	user4 := user3
	fmt.Printf("before user3:%v, user4:%v\n", user3, user4)
	user4.Name = "update"
	fmt.Printf("after user3:%v, user4:%v\n", user3, user4)

	fmt.Println("引用类型, DeepCopy深拷贝:")
	user5 := &User{
		Name: "why",
	}
	user6 := new(User)
	fmt.Printf("before user5:%v, user6:%v\n", user5, user6)
	DeepCopy(user6, user5)
	user6.Name = "update"
	fmt.Printf("after user5:%v, user6:%v\n", user5, user6)

	fmt.Println("值类型, DeepCopy深拷贝:")
	user7 := User{
		Name: "why",
	}
	user8 := new(User)
	fmt.Printf("before user7:%v, user8:%v\n", user7, user8)
	DeepCopy(user8, user7)
	user8.Name = "update"
	fmt.Printf("after user7:%v, user8:%v\n", user7, user8)

}
