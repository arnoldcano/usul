package sandbox

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func saveFile(language, code string) (*os.File, error) {
	file, err := os.Create(fmt.Sprintf("%s/%s", FilePath, getFileName(language)))
	if err != nil {
		return nil, err
	}
	_, err = file.WriteString(code)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func getFileName(language string) string {
	chars := "abcdefghijklmnopqrstuvwxyz0123456789"
	rand.Seed(time.Now().UTC().UnixNano())
	tempName := make([]byte, 8)
	for i := 0; i < 8; i++ {
		tempName[i] = chars[rand.Intn(len(chars))]
	}
	fileName := fmt.Sprintf("%s.%s", string(tempName), getFileExtension(language))
	return fileName
}

func getFileExtension(language string) string {
	var ext string
	switch language {
	case "ruby":
		ext = "rb"
	case "python":
		ext = "py"
	case "nodejs":
		ext = "js"
	}
	return ext
}
