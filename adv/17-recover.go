package main 

// 无效一
// recover捕获的是祖父级调用时的异常，直接调用时无效
/*
func main() {
	recover() 
	panic(1)
}
*/

// 无效二
//直接defer调用也是无效
/*
func main() {
	defer recover()
	panic(1)
}
*/


// 无效三
// defer调用时多层嵌套依然无效：
// func main() {
// 	defer func() {
// 		func() {
// 			recover()
// 		}()
// 	}()
// 	panic(1)
// }

// 必须在defer函数中直接调用才有效
func main() {
	defer func() {
		recover()
	}()
	panic(1)
}
