// run cmd:  go test -bench=. 07_large_slice_range_test.go
package test

import "testing"

func Test(t *testing.T) {
}

func CreateBigSlice(count int) [][4096]int {
	ret := make([][4096]int, count)
	for i := 0; i < count; i++ {
		ret[i] = [4096]int{}
	}
	return ret
}

func BenchmarkRangeHiPerformace(b *testing.B) {
	v := CreateBigSlice(1 << 12)

	for i := 0; i < b.N; i++ {
		l := len(v)
		var tmp [4096]int
		for k := 0; k < l; k++ {
			tmp = v[k]
		}
		_ = tmp
	}
}

func BenchmarkRangeLowPerformace(b *testing.B) {
	v := CreateBigSlice(1 << 12)
	for i := 0; i < b.N; i++ {
		var tmp [4096]int
		for _, e := range v {
			tmp = e
		}
		_ = tmp
	}
}
