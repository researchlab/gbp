package main

import "fmt"
/*
声明一个空slice，默认len和cap的值都为0。声明长度为0的array的方式: var b [0]int，这两者声明方式非常相似
slice在新增元素的时候，如果len超过cap，会动态扩容cap，

结论：s = s[low : high : max] 切⽚的三个参数的切⽚截取的意义为 low为截取的起始下标（含）， high为窃取的结束下标（不含high），max为
切⽚保留的原切⽚的最⼤下标（不含max）；即新切⽚从⽼切⽚的low下标元素开始，len = high - low, cap = max - low；high 和 max⼀旦超出在⽼
切⽚中越界，就会发⽣runtime err，slice out of range。另外如果省略第三个参数的时候，第三个参数默认和第⼆个参数相同，即len = cap。
*/
func main() {
	//array := [...]int{1, 2, 3, 4, 5}  这样声明的是[5]int 数组，形参也要写成 [5]int 否则没法用
	array := []int{1, 2, 3, 4, 5}
	s1 := array[:2:4] // 指定cap 容量
	s2 := array[2:]
	show := func(key string, array []int) {
		fmt.Printf("%s:%v  len(%s):%v cap(%s):%v\n", key, array, key, len(array), key, cap(array))
	}
	show("array", array)
	show("s1", s1)
	show("s2", s2)
}

/**
output 
array:[1 2 3 4 5]  len(array):5 cap(array):5
s1:[1 2]  len(s1):2 cap(s1):4
s2:[3 4 5]  len(s2):3 cap(s2):3
*/
