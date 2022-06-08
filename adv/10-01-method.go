package main

import "fmt"

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {}

func live() People {
	var stu *Student
	return stu
}

type Null interface{}

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBB")
	}
	// output BBBBBB  因为这个interface 带有方法集

	var n *Null
	if n == nil {
		fmt.Println("CCCCCCC")
	} else {
		fmt.Println("DDDDDDDD")
	}
	// output CCCCCCC  因为是空的 interface
}

/*
考点：interface内部结构

解答：很经典的题！这个考点是很多人忽略的interface内部结构。 go中的接口分为两种一种是空的接口类似这样：

var in interface{}
另一种如题目：

type People interface {
    Show()
}
他们的底层结构如下：

type eface struct {
    //空接口
    _type *_type         //类型信息
    data  unsafe.Pointer //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
}
type iface struct {
    //带有方法的接口
    tab  *itab           //存储type信息还有结构实现方法的集合
    data unsafe.Pointer  //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)}
type _type struct {
    size       uintptr  //类型大小
    ptrdata    uintptr  //前缀持有所有指针的内存大小
    hash       uint32   //数据hash值
    tflag      tflag    align      uint8    //对齐
    fieldalign uint8    //嵌入结构体时的对齐
    kind       uint8    //kind 有些枚举值kind等于0是无效的
    alg        *typeAlg //函数指针数组，类型实现的所有方法
    gcdata    *byte    str       nameOff    ptrToThis typeOff
}
type itab struct {
    inter  *interfacetype  //接口类型
    _type  *_type          //结构类型
    link   *itab    bad    int32    inhash int32    fun    [1]uintptr      //可变大小 方法集合
}
可以看出iface比eface 中间多了一层itab结构。 itab 存储_type信息和[]fun方法集，从上面的结构我们就可得出，
因为data指向了nil 并不代表interface 是nil，所以返回值并不为空，这里的fun(方法集)定义了接口的接收规则，
在编译的过程中需要验证是否实现接口结果：

BBBBBBB
*/
