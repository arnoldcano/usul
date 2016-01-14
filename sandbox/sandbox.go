package sandbox

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	ContainerName    = "usul"
	ContainerBinPath = "/usr/bin"
	FilesDir         = "files"
	Timeout          = 5
)

type CompileCommand struct {
	Language string `json:"language"`
	Code     string `json:"code"`
	Output   string `json:"output"`
}

func (c *CompileCommand) Run() error {
	var output bytes.Buffer

	file, err := c.getTempFile()
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())

	opts, err := c.getOptions(file.Name())
	if err != nil {
		return err
	}

	log.Printf("Running 'docker %s'\n", strings.Join(opts, " "))

	command := exec.Command("docker", opts...)
	command.Stdout = &output
	command.Stderr = &output
	if err := command.Start(); err != nil {
		return err
	}
	timer := time.AfterFunc(Timeout*time.Second, func() {
		log.Printf("Killing 'docker %s'\n", strings.Join(opts, " "))
		command.Process.Kill()
	})
	command.Wait()
	timer.Stop()

	c.Output = output.String()

	return nil
}

func (c *CompileCommand) getTempFile() (*os.File, error) {
	file, err := os.Create(fmt.Sprintf("%s/%s", FilesDir, c.getFileName()))
	if err != nil {
		return nil, err
	}

	_, err = file.WriteString(c.Code)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (c *CompileCommand) getOptions(fileName string) ([]string, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	filesPath := fmt.Sprintf("%s/%s", path, FilesDir)
	opts := []string{
		"run",
		"--rm",
		"-v",
		fmt.Sprintf("%s:/%s", filesPath, FilesDir),
		ContainerName,
		fmt.Sprintf("%s/%s", ContainerBinPath, c.Language),
		fmt.Sprintf("/%s", fileName),
	}

	return opts, nil
}

func (c *CompileCommand) getFileName() string {
	rand.Seed(time.Now().UTC().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyz0123456789"
	tempName := make([]byte, 8)
	for i := 0; i < 8; i++ {
		tempName[i] = chars[rand.Intn(len(chars))]
	}
	fileName := fmt.Sprintf("%s.%s", string(tempName), c.getFileExtension())

	return fileName
}

func (c *CompileCommand) getFileExtension() string {
	var ext string

	switch c.Language {
	case "ruby":
		ext = "rb"
	case "python":
		ext = "py"
	case "nodejs":
		ext = "js"
	}

	return ext
}
