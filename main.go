package main

import (
	"log"
	"net/http"

	"github.com/arnoldcano/usul/runner"
)

func main() {
	http.HandleFunc("/run", runner.RunHandler)

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
