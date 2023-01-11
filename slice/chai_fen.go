package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var arr, arr1 []string
	for i := 0; i < 11; i++ {
		str := strconv.Itoa(i)
		arr = append(arr, str)
		arr1 = append(arr1, fmt.Sprintf("%s-arr1", str))
	}
	fmt.Println("println:", group(arr, 5))
	res := group(arr1, 5)
	arr1[0] = "111"
	for _, v := range res {
		rt := strings.Join(v, ",")
		fmt.Println(rt)
	}
}

func group(in []string, subGroupLength int64) [][]string {
	max := int64(len(in))
	segmens := make([][]string, 0)
	quotient := max / subGroupLength
	remainder := max % subGroupLength
	safetyCopy := func(in []string)[]string{
		tmp := make([]string, len(in))
		copy(tmp, in)
		return tmp
	}
	i := int64(0)
	for i=int64(0); i < quotient ; i++ {
		segmens = append(segmens, safetyCopy(in[i*subGroupLength:(i+1)*subGroupLength]))
	}
	if quotient == 0 || remainder != 0 {
		segmens = append(segmens, safetyCopy(in[i*subGroupLength:i*subGroupLength+remainder]))
	}
	return segmens
}
