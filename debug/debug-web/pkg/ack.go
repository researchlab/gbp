package pkg

import (
	"fmt"
	"net/http"

	"github.com/researchlab/gbp/debug/debug-web/utils"
)

func Ack(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println("path:", path)
	fmt.Fprintf(w, "%s", "world")

	utils.Show()
}
