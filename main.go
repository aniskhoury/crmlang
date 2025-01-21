package main

import (
	/*
		"context"
		"database/sql"
		"fmt"
		"html/template"

		_ "github.com/go-sql-driver/mysql"
		"github.com/gorilla/sessions"*/

	"crmlang/configuration"
	"crmlang/controllers/index"
	"net/http"
)

var (
	contador int = 0
)

func main() {

	fileServer := http.FileServer(http.Dir("resources"))
	http.Handle("/resources/", http.StripPrefix("/resources", fileServer))
	http.HandleFunc("/", index.IndexFunc)
	http.ListenAndServe(":"+configuration.PortHTTPServer, nil)

}

//go run ./
