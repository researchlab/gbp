package utils

import (
	"fmt"
	"net/http"
)

func Ack(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println("path:", path)
	fmt.Fprintf(w, "%s", "world")
}
