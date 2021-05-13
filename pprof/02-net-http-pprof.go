package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func counter() {
	slice := make([]int, 0)
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 1 + 2 + 3 + 4 + 5
		slice = append(slice, c)
	}
}

func workForever() {
	for {
		go counter()
		time.Sleep(time.Second)
	}
}

func httpGet(w http.ResponseWriter, r *http.Request) {
	counter()
}

func main() {
	go workForever()
	http.HandleFunc("/get", httpGet)
	http.ListenAndServe("localhost:8000", nil)
}
