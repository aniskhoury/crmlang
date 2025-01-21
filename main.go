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
	handlefunc "crmlang/handleFunc"
	"net/http"
)

func main() {

	handlefunc.CallAllHandleFunctions()
	http.ListenAndServe(":"+configuration.PortHTTPServer, nil)

}

//go run ./
