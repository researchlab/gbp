package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hi", hi)
	http.ListenAndServe(":8081", nil)
}

func hi(w http.ResponseWriter, r *http.Request) {
	hostName, _ := os.Hostname()
	fmt.Fprintf(w, "Hostname: %s", hostName)
}
