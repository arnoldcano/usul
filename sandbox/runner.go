package sandbox

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func runFile(language, fileName string) ([]byte, error) {
	var output bytes.Buffer
	opts, err := getOptions(language, fileName)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command("docker", opts...)
	log.Printf("Running 'docker %s'\n", strings.Join(opts, " "))
	cmd.Stdout = &output
	cmd.Stderr = &output
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	timer := time.AfterFunc(Timeout*time.Second, func() {
		log.Printf("Killing 'docker %s'\n", strings.Join(opts, " "))
		cmd.Process.Kill()
	})
	cmd.Wait()
	timer.Stop()
	return output.Bytes(), nil
}

func getOptions(language, fileName string) ([]string, error) {
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
		"usul",
		fmt.Sprintf("%s/%s", BinDir, language),
		fmt.Sprintf("/%s", fileName),
	}
	return opts, nil
}
