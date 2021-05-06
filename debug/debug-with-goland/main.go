package main

import "time"

func main() {

	c := 0
	for {
		time.Sleep(time.Second * 1)
		c++

		println(c)
	}
}
