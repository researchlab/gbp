package main

import "fmt"

type student struct {
	Name string
	Age  int
}

var stus = []student{
	{
		Name: "zhou",
		Age:  10,
	},
	{
		Name: "li",
		Age:  11,
	},
	{
		Name: "wang",
		Age:  120,
	},
}

func error_copy_student() map[string]*student {
	m := make(map[string]*student)
	for _, stu := range stus {
		m[stu.Name] = &stu // 这里&stu 这个指针是一直指向同一个指针, 直到for结束，这个指针就是最后一个元素
	}
	return m
}

func error_modify_age() {
	for _, stu := range stus {
		stu.Age = stu.Age + 10 // 这个修改只反应在临时变量 stu上，并没有修改stus
	}
	fmt.Println(stus)
	/*
		[{zhou 10} {li 11} {wang 120}]
	*/
}

func right_copy_student() map[string]*student {
	m := make(map[string]*student)
	// for i, stu := range stus {
	// 	m[stu.Name] = &stus[i] // 注意这里是浅拷贝， 这里改变 也会影响到stus
	// }
	// 浅拷贝
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i] // 注意这里是浅拷贝， 这里改变 也会影响到stus
	}

	// 深拷贝
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &student{
			Name: stus[i].Name,
			Age:  stus[i].Age,
		}
	}

	return m
}
func show(in map[string]*student) {
	for k, v := range in {
		fmt.Println(k, "=>", v)
	}
	fmt.Println("================")
}

func main() {
	show(error_copy_student())
	/*
		output
		zhou => &{wang 120}
		li => &{wang 120}
		wang => &{wang 120}
	*/
	m := right_copy_student()
	li := m["li"]
	li.Age = 100
	li.Name = "li update"
	m["li"] = li

	show(m)
	fmt.Println(stus)
	error_modify_age()

}

/*
for 拷贝时，要特别注意 浅拷贝和深拷贝

不带&实例化的结构体，是值类型，赋值是深拷贝。
& 和 new 实例化的结构体，都是引用类型，赋值都是浅拷贝。
无论是引用类型还是值类型，通过序列化 - 反序列化的方式，都可以转为深拷贝。

*/
