package utils

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Info struct {
	Num int `json:"num"`
}

var info *Info

var once sync.Once

func lazzy() *Info{
	once.Do(func() {
		info = &Info{Num: 1}
	})
	return info
}

func Show() {

	fmt.Println((lazzy()).Num)
	b, _ := json.Marshal(info)
	fmt.Println(string(b))
}
