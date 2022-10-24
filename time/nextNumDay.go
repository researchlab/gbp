package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	year, month, day := now.Date()
	fmt.Println(year, month, day)
	fmt.Println(year, uint(month), day)
// year month, day
	lastTime := now.AddDate(0, 0, -7)
	lastYear, lastMonth, lastDay := lastTime.Date()
	fmt.Println(lastYear, int(lastMonth), lastDay)
}
