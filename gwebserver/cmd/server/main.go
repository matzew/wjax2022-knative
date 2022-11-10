package main

import (
	"log"
	"net/http"
	"time"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	// Simulate at least a bit of processing time.
	time.Sleep(100 * time.Millisecond)
	log.Print("Got request for processing...")

	// We abort the request and tell client that it
	// was semantically not correct...
	w.Header().Add("kne-workerprocessor", "Semantic error with request")
	w.Header().Add("kne-workerhints", "Patch your data model")
	w.WriteHeader(http.StatusUnprocessableEntity)

	// ship it!
	w.Write([]byte("Unable to process request!"))
}

func main() {

	m := http.NewServeMux()
	m.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", m)
}
