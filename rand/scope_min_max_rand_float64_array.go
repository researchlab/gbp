package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(randFloats(1.10, 101.99, 5))
	fmt.Println(randFloats(1.10, 101.99, 5))
}
