package main

import (
	"fmt"
	"time"
)

func diffDays(start, end int64) int {
	startTime := time.Unix(start, 0)
	endTime := time.Unix(end, 0)
	sub := int(endTime.Sub(startTime).Hours())
	days := sub / 24
	if (sub % 24) > 0 {
		days = days + 1
	}
	return days
}

func main() {
	fmt.Println("should be 0:", diffDays(time.Now().Unix(), time.Now().Unix()))
	fmt.Println("should be 1:", diffDays(time.Now().Unix(), time.Now().Unix()+int64(24*60*60)))
	fmt.Println("should be 2:", diffDays(time.Now().Unix(), time.Now().Unix()+int64(24*60*60+1)))
	sub := time.Unix(time.Now().Unix()+int64(24*60*60),0).Sub(time.Now()).Hours()
	fmt.Println(sub)
	fmt.Println(sub/24)
	//fmt.Println(sub % 24)
	t := time.Now().Unix()
	tt := t+int64(24*60*60+1)
	num := (tt -t)/86400
	delta := (tt-t)%86400
	fmt.Println(num)
	fmt.Println(delta)
}
