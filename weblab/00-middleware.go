package main

import "net/http"

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8881", nil)
}
