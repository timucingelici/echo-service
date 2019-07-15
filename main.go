package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", echoHandler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {

}