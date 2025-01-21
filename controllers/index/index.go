package index

import (
	"fmt"
	"net/http"
)

var (
	contador int = 0
)

func IndexFunc(w http.ResponseWriter, r *http.Request) {
	contador = contador + 1
	fmt.Fprintf(w, "hola mundo %d", contador)
}
