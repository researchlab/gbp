package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a[:len(a)-2]

	c := append(a, []int{4, 5}...)
	d := append(a, []int{7, 8}...)

	e := append(a[:len(a)-2], []int{4, 5}...)
	f := append(a[:len(a)-2], []int{7, 8}...)

	g := append(a[:len(a)-2], []int{4, 5,6}...)
	h := append(a[:len(a)-2], []int{7, 8,9}...)


	fmt.Println("a:", a,  " len:", len(a), " cap:", cap(a))
	fmt.Println("b:", b,  " len:", len(b), " cap:", cap(b))
	fmt.Println("c:", c,  " len:", len(c), " cap:", cap(c))
	fmt.Println("d:", d,  " len:", len(d), " cap:", cap(d))
	fmt.Println("e:", e,  " len:", len(e), " cap:", cap(e))
	fmt.Println("f:", f,  " len:", len(f), " cap:", cap(f))
	fmt.Println("g:", g,  " len:", len(g), " cap:", cap(g))
	fmt.Println("h:", h,  " len:", len(h), " cap:", cap(h))
}

// output 
// a: [1 7 8]  len: 3  cap: 3
// b: [1]  len: 1  cap: 3
// c: [1 2 3 4 5]  len: 5  cap: 6
// d: [1 2 3 7 8]  len: 5  cap: 6
// e: [1 7 8]  len: 3  cap: 3
// f: [1 7 8]  len: 3  cap: 3
// g: [1 4 5 6]  len: 4  cap: 6
// h: [1 7 8 9]  len: 4  cap: 6
//
