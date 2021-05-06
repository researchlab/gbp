package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/hi", hi)
	greet := "Hello World!"
	fmt.Println("greet:", greet)
	func01()
	func02()
	http.HandleFunc("/sum", sum)
	fmt.Println("running on port:8081")
	http.ListenAndServe(":8081", nil)
}

func func01() {
	fmt.Println("func01")
}

func func02() {
	fmt.Println("func02")
}

func sum(w http.ResponseWriter, r *http.Request) {
	result := 0
	if r.Method == "GET" {
		vals := r.URL.Query()
		_num := vals.Get("num")
		if len(_num) != 0 {
			num, _ := strconv.Atoi(_num)
			if num != 0 {
				for i := 0; i <= num; i++ {
					result += i
				}
			}
		}
	}
	fmt.Fprintf(w, "result: %d\n", result)
}
func hi(w http.ResponseWriter, r *http.Request) {
	doHi(w)
}

func doHi(w http.ResponseWriter) {
	hostName, _ := os.Hostname()
	fmt.Fprintf(w, "Hostname: %s\n", hostName)
}
