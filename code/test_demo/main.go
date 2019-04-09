package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Add(value1 int, value2 int) int {
	return value1 + value2
}

func Subtract(value1 int, value2 int) int {
	return value1 - value2
}

func RootEndPoint(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	response.Write([]byte("Hello world"))
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndPoint).Methods("GET")
	log.Fatal(http.ListenAndServe("1234", router))
}
