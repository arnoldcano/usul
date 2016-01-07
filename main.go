package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type command struct {
	name     string
	language string
	file     string
}

func compileHandler(w http.ResponseWriter, r *http.Request) {
	language := r.FormValue("language")
	file := r.FormValue("file")

	output := make(chan []byte)
	command := &command{
		name:     "docker",
		language: language,
		file:     file,
	}
	go run(command, output)
	out := <-output
	close(output)

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func run(command *command, output chan []byte) {
	opts := getCommandOptions(command)
	fmt.Printf("usul> docker %s\n", strings.Join(opts, " "))
	out, _ := exec.Command(command.name, opts...).CombinedOutput()
	output <- out
}

func getCommandOptions(command *command) []string {
	opts := []string{
		"run",
		"--rm",
		"-v",
		fmt.Sprintf("%s:/code", getPath()),
		"usul",
	}
	opts = getCompilerOptions(command, opts)

	return opts
}

func getCompilerOptions(command *command, opts []string) []string {
	switch command.language {
	case "go":
		opts = append(opts, fmt.Sprintf("/usr/bin/%s", "go"))
		opts = append(opts, fmt.Sprintf("run"))
	case "ruby":
		opts = append(opts, fmt.Sprintf("/usr/bin/%s", "ruby"))
	case "python":
		opts = append(opts, fmt.Sprintf("/usr/bin/%s", "python"))
	}
	opts = append(opts, fmt.Sprintf("/code/%s", command.file))

	return opts
}

func getPath() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return path
}

func main() {
	http.HandleFunc("/compile", compileHandler)

	fmt.Println("Listing on port 8080...")
	http.ListenAndServe(":8080", nil)
}
