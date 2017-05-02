// Simple Golang program to test new multi-stage builds feature in Docker
package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port     = ":8080"
	greeting = "Hello Golang World!"
)

func main() {
	log.Println("Listening on port ", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, greeting)
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
