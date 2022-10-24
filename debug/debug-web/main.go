package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/researchlab/gbp/debug/debug-web/pkg"
)

func main() {
	http.HandleFunc("/hello", pkg.Ack)
	fmt.Println("service start at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
