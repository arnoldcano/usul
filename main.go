package main

import (
	"fmt"
	"net/http"

	"github.com/arnoldcano/usul/server"
)

func main() {
	http.HandleFunc("/compile", server.CompileHandler)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
