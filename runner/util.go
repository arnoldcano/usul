package runner

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func runFile(lang string, f *os.File) (string, error) {
	var b bytes.Buffer

	c := exec.Command(lang, []string{f.Name()}...)
	c.Stdout = &b
	c.Stderr = &b
	if err := c.Start(); err != nil {
		return "", err
	}
	log.Printf("Running temp file %s", f.Name())
	t := time.AfterFunc(Timeout*time.Second, func() {
		log.Printf("Process %d timed out", c.Process.Pid)
		c.Process.Kill()
	})
	c.Wait()
	t.Stop()

	return b.String(), nil
}

func saveTempFile(lang, code string) (*os.File, error) {
	d, err := ioutil.TempDir("", "")
	if err != nil {
		return nil, err
	}

	f, err := os.Create(fmt.Sprintf("%s/%s", d, getFileName(lang)))
	if err != nil {
		return nil, err
	}

	_, err = f.WriteString(code)
	if err != nil {
		return nil, err
	}
	log.Printf("Saved temp file %s", f.Name())

	return f, nil
}

func getFileName(lang string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	c := "abcdefghijklmnopqrstuvwxyz0123456789"
	t := make([]byte, 8)
	for i := 0; i < 8; i++ {
		t[i] = c[rand.Intn(len(c))]
	}
	n := fmt.Sprintf("%s.%s", string(t), getFileExtension(lang))

	return n
}

func getFileExtension(lang string) string {
	var e string

	switch lang {
	case "ruby":
		e = "rb"
	case "python":
		e = "py"
	case "js":
		e = "js"
	}

	return e
}
