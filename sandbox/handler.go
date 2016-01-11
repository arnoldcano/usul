package sandbox

import (
	"net/http"
	"os"
)

const (
	BinDir   = "/usr/bin"
	FilesDir = "files"
	Timeout  = 5
)

func RunHandler(w http.ResponseWriter, r *http.Request) {
	language := r.FormValue("language")
	code := r.FormValue("code")
	file, err := saveFile(language, code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	defer os.Remove(file.Name())
	output, err := runFile(language, file.Name())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
