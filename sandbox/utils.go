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

func saveTempFile(lang, code string) (*os.File, error) {
	file, err := getTempFile(lang)
	if err != nil {
		return nil, err
	}
	_, err = file.WriteString(code)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func getTempFile(lang string) (*os.File, error) {
	rand.Seed(time.Now().UTC().UnixNano())
	temp := make([]byte, 8)
	for i := 0; i < 8; i++ {
		temp[i] = Chars[rand.Intn(len(Chars))]
	}
	fileName := fmt.Sprintf("%s.%s", string(temp), getExtension(lang))
	file, err := os.Create(fmt.Sprintf("%s/%s", FilePath, fileName))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func getExtension(lang string) string {
	var ext string
	switch lang {
	case "ruby":
		ext = "rb"
	case "python":
		ext = "py"
	case "nodejs":
		ext = "js"
	}
	return ext
}

func compileTempFile(lang string, file *os.File) ([]byte, error) {
	var output bytes.Buffer
	opts, err := getOptions(lang, file)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command("docker", opts...)
	log.Printf("Running command 'docker %s'\n", strings.Join(opts, " "))
	cmd.Stdout = &output
	cmd.Stderr = &output
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	timer := time.AfterFunc(Timeout*time.Second, func() {
		log.Printf("Killing command 'docker %s'\n", strings.Join(opts, " "))
		cmd.Process.Kill()
	})
	cmd.Wait()
	timer.Stop()
	return output.Bytes(), nil
}

func getOptions(lang string, file *os.File) ([]string, error) {
	path, err := getPath()
	if err != nil {
		return nil, err
	}
	opts := []string{
		"run",
		"--rm",
		"-v",
		fmt.Sprintf("%s/%s:/%s", path, FilePath, FilePath),
		ContainerName,
		fmt.Sprintf("%s/%s", CompilerPath, lang),
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
