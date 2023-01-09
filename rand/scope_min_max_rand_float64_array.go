package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"
	RR "crypto/rand"
)

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

// 生成区间[-m, n]的安全随机数
func RangeRand(min, max int64) int64 {
	if min > max {
		panic("the min is greater than max")
	}
	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := RR.Int(RR.Reader, big.NewInt(max+1+i64Min))
		return result.Int64() - i64Min
	} else {
		result, _ := RR.Int(RR.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(randFloats(1.10, 101.99, 5))
	fmt.Println(randFloats(1.10, 101.99, 5))
	fmt.Println(RangeRand(-100, 100))
}
