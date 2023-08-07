package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

//go.sum checks whether the downloaded package is original it does this by comparing the hash and verify its
//All the lib which we pulled by go get is stored in go/pkg/mod/cache

func main() {
	fmt.Println("mod  in Golang")
	greet()

	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))

}

func greet() {
	fmt.Println("Welcome to Mux framework")
}

func serverHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcome to Mux</h1>"))

}
