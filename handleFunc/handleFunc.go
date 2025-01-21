package handlefunc

import (
	"crmlang/controllers/index"
	"net/http"
)

func CallAllHandleFunctions() {
	fileServer := http.FileServer(http.Dir("resources"))
	http.Handle("/resources/", http.StripPrefix("/resources", fileServer))
	http.HandleFunc("/", index.IndexFunc)
}
