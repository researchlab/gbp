package main

import "math/rand"

func main() {
	sum := 0
	for {
		n := rand.Intn(1e6)
		sum += n
		if sum%42 == 0 {
			panic(":(")
		}
	}
}
