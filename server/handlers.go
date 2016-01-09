package server

import (
	"net/http"
	"os"
)

const (
	ContainerName = "usul"
	CompilerPath  = "/usr/bin"
	FilePath      = "sandbox"
	Chars         = "abcdefghijklmnopqrstuvwxyz0123456789"
	Timeout       = 5
)

func CompileHandler(w http.ResponseWriter, r *http.Request) {
	lang := r.FormValue("lang")
	code := r.FormValue("code")
	file, err := saveTempFile(lang, code)
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())
	output, err := compileTempFile(lang, file)
	if err != nil {
		panic(err)
	}
	w.Write(output)
}
