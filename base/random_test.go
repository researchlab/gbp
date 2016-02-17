package base

import (
	"testing"
)

func Test_GenerateRandomNumber(t *testing.T) {
	for i := 0; i < 5; i++ {
		nums := GenerateRandomNumber(10, 30, 10)
		_, err := IsDuplicate(nums)
		if err == nil {
			panic(err.Error())
		}
	}
}
