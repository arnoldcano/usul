package runner

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const Timeout = 5

type Request struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

type Response struct {
	Output string `json:"output"`
}

func RunHandler(w http.ResponseWriter, r *http.Request) {
	var r2 Request
	var w2 Response

	err := json.NewDecoder(r.Body).Decode(&r2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Received request from %s", r.UserAgent())

	f, err := saveTempFile(r2.Language, r2.Code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(filepath.Dir(f.Name()))

	w2.Output, err = runFile(r2.Language, f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(w2); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Printf("Sent response to %s", r.UserAgent())
	log.Printf("Removed temp dir %s", filepath.Dir(f.Name()))
}
