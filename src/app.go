package main

import "net/http"

func main() {
	http.HandleFunc("/hello-world", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello world"))
	})
	http.ListenAndServe(":3000", nil)
}
