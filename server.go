package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func healthCheck(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Up and running...")

}

func main() {
	router := mux.NewRouter()
	const port string = ":8000"

	router.HandleFunc("/", healthCheck)

	router.HandleFunc("/posts", getPosts).Methods("GET")

	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))

}
