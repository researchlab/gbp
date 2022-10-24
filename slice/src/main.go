package main

import (
	"encoding/json"
	"fmt"
	"pkgs"
)

func main() {
	in := pkgs.MakeInfo()

	bin, _ := json.Marshal(in)
	fmt.Println(string(bin))

	m := make(map[int]bool)
	//for _, n := range []int{0,1}{
	//	m[n]= true
	//}
	fmt.Println(m)
	j := 0
	for _, v := range *in {
		if _, ok := m[v.Num]; ok {
			//in = append(in[:idx], in[idx+1:]...)
			(*in)[j] = v
			j++
		}
	}
	*in = (*in)[:j]
	bb, _ := json.Marshal(in)
	fmt.Println(string(bb))
	fmt.Println(j)
}
