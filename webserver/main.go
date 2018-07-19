package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello world</h1>")
}
