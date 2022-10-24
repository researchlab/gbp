package main 

/*
#include <stdio.h>

static void SayHello(const char* s) {
	puts(s);
}
*/
import "C"

// 注意 上面的函数和 import "C"要紧挨着，不能有空行 否则报错
func main(){
	C.SayHello(C.CString("Hello, World\n"))
}
