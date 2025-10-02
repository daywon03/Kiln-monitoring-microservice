package main

import (
	"log"
    "net/http"
)

func main() {
	log.Println("Starting server on port 8080")
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	log.Println("listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}