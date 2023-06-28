package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

//go:embed css
var FS embed.FS

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	http.Handle("/", templ.Handler(home()))
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.FS(FS))))
	return http.ListenAndServe("localhost:8081", nil)
}
