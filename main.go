package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	containerName = "usul"
	compilerPath  = "/usr/bin"
	filePath      = "code"
	chars         = "abcdefghijklmnopqrstuvwxyz0123456789"
)

func compileHandler(w http.ResponseWriter, r *http.Request) {
	lang := r.FormValue("lang")
	code := r.FormValue("code")

	file, err := saveToFile(lang, code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer os.Remove(file.Name())

	output := make(chan []byte)
	go run(lang, file, output)
	out := <-output
	close(output)

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func saveToFile(lang, code string) (*os.File, error) {
	name := getFileName(lang)

	file, err := os.Create(fmt.Sprintf("%s/%s", filePath, name))
	if err != nil {
		return nil, err
	}

	_, err = file.WriteString(code)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func getFileName(lang string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	bytes := make([]byte, 8)
	for i := 0; i < 8; i++ {
		bytes[i] = chars[rand.Intn(len(chars))]
	}
	name := string(bytes) + getExtension(lang)

	return name
}

func getExtension(lang string) string {
	var ext string
	switch lang {
	case "ruby":
		ext = ".rb"
	case "python":
		ext = ".py"
	case "nodejs":
		ext = ".js"
	}

	return ext
}

func run(lang string, file *os.File, output chan []byte) error {
	opts, err := getCommandOptions(lang, file)
	if err != nil {
		return err
	}

	fmt.Printf("%s> docker %s\n", containerName, strings.Join(opts, " "))

	out, err := exec.Command("docker", opts...).CombinedOutput()
	if err != nil {
		return err
	}
	output <- out

	return nil
}

func getCommandOptions(lang string, file *os.File) ([]string, error) {
	path, err := getPath()
	if err != nil {
		return nil, err
	}

	opts := []string{
		"run",
		"--rm",
		"-v",
		fmt.Sprintf("%s/%s:/%s", path, filePath, filePath),
		containerName,
		fmt.Sprintf("%s/%s", compilerPath, lang),
		fmt.Sprintf("/%s", file.Name()),
	}

	return opts, nil
}

func getPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return path, nil
}

func main() {
	http.HandleFunc("/compile", compileHandler)

	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
