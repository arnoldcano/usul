package sandbox

import (
	"encoding/json"
	"net/http"
)

func CompileHandler(w http.ResponseWriter, r *http.Request) {
	command := NewCompileCommand()

	err := json.NewDecoder(r.Body).Decode(command)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := command.Run(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(command); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
