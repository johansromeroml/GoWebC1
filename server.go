package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", ping)

	mux.HandleFunc("/greetings", greetings)

	http.ListenAndServe(":8080", mux)
}

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong")
}

type Person struct {
	FirstName string
	LastName  string
}

func greetings(w http.ResponseWriter, req *http.Request) {
	var per Person
	payload, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if err := json.Unmarshal(payload, &per); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Hello %s %s", per.FirstName, per.LastName)
}
