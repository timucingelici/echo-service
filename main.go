package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", echoHandler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {

	// Read the body from the request
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error reading the body: %v", err)
		http.Error(w, "Error reading the body", http.StatusBadRequest)
	}

	// Get the length of the body
	length := len(body)

	// Print the info
	log.Printf(
		"Got a %s request with the body size of %d bytes, so I'll mirror the request and send %d bytes back",
		r.Method, length, length,
	)

	// Write the body back
	_, err = w.Write(body)

	if err != nil {
		log.Printf("Error writing the body")
		http.Error(w, "Error writing the body", http.StatusInternalServerError)
	}
}