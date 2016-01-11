package main

import (
	"fmt"
	"net/http"

	"github.com/arnoldcano/usul/sandbox"
)

func main() {
	http.HandleFunc("/run", sandbox.RunHandler)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
