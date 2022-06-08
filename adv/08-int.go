package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		// fmt.Println("i" + string(i)) // output ii
		fmt.Println(fmt.Sprintf("i%d", i))
	}
}

// 注意， 数字不能通过string(数字) 强转
