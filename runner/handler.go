package runner

import (
	"encoding/json"
	"log"
	"net/http"
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
		writeError(w, err)
		return
	}
	log.Printf("Received request from %s", r.RemoteAddr)

	f, err := saveTempFile(r2.Language, r2.Code)
	if err != nil {
		writeError(w, err)
		return
	}
	defer removeTempFile(f)

	w2.Output, err = runFile(r2.Language, f)
	if err != nil {
		writeError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(w2); err != nil {
		writeError(w, err)
	}
	log.Printf("Sent response to %s", r.RemoteAddr)
}

func writeError(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		panic(err)
	}
}
