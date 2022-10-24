package main 

//#include <stdio.h>
import "C"
// 通过import "C"语句启用CGO特性，同时包含C语言的<stdio.h>头文件。然后通过CGO包的C.CString函数将Go语言字符串转为C语言字符串，最后调用CGO包的C.puts函数向标准输出窗口打印转换后的C字符串。
// 注意: 没有在程序退出前释放C.CString创建的C语言字符串；
// 没有释放使用C.CString创建的C语言字符串会导致内存泄漏。但是对于这个小程序来说，这样是没有问题的，因为程序退出后操作系统会自动回收程序的所有资源。
func main() {
	C.puts(C.CString("Hello, World\n"))
}
