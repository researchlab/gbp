package main

import "fmt"

func main() {
	var s []int = make([]int, 5)
	fmt.Println("len:", len(s), " cap:", cap(s))
	// error  index out of  range 5
	// s[100]++
	// fmt.Println("len:", len(s), " cap:", cap(s))
	// error
	result := findRepeatNumber2([]int{1,2,100,100})
	fmt.Println("result:",result)
}

func findRepeatNumber(nums []int) int {
	var a []int = make([]int, len(nums))
	var ans int = 0
	for _, v := range nums {
		a[v]++
		if a[v] >= 2 {
			ans = v
			break
		}
	}
	return ans
}

func findRepeatNumber2(nums []int)int{
	var a []int  = make([]int, len(nums))
	var ant int = 0
	for _, v := range nums {
		if a[v] == 1 {
			ant = v 
			break
		}
		a[v] = 1
	}
	return ant 
}
