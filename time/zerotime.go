package main

import (
	"fmt"
	"time"
)

func main() {
	data := map[string][]int{
		"今天0点":    []int{0, 0, 0},
		"昨天0点":    []int{0, 0, -1},
		"上一周0点":   []int{0, 0, -7},
		"上一个月0点":  []int{0, -1, 0},
		"上三个月0点":  []int{0, -3, 0},
		"上六个月0点":  []int{0, -6, 0},
		"上12个月0点": []int{0, -12, 0},
		"下一周0点":   []int{0, 0, 7},
		"下一个月0点":  []int{0, 1, 0},
		"下三个月0点":  []int{0, 3, 0},
		"下六个月0点":  []int{0, 6, 0},
		"下12个月0点": []int{0, 12, 0},
	}
	for k, v := range data {
		fmt.Println(k)
		ts := zeroTimestamp(v[0], v[1], v[2])
		fmt.Println(ts, format(ts))
	}
}

func format(ts int64) string {
	return time.Unix(ts, 0).Format(time.RFC3339)
}
func zeroTimestamp(year, month, day int) int64 {
	ts := time.Now().AddDate(year, month, day)
	return time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, ts.Location()).Unix()
}
