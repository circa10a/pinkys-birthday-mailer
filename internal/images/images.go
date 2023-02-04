package images

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
)

// ErrNoImagesFound is returned when no images are found in the provided directory.
var ErrNoImagesFound = errors.New("no images found")

// Find gets all of the images in a directory.
func Find(dir string) ([]string, error) {
	fileNames := []string{}

	files, err := os.ReadDir(dir)
	if err != nil {
		return fileNames, err
	}

	for _, file := range files {
		filePath := path.Join(dir, file.Name())
		if isValidImage(filePath) {
			fileNames = append(fileNames, filePath)
		}
	}

	if len(fileNames) == 0 {
		return fileNames, ErrNoImagesFound
	}

	return fileNames, nil
}

// isValidImage determines if a file path is a jpeg or png.
func isValidImage(fileName string) bool {
	buf := make([]byte, 512)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	defer file.Close()

	_, err = file.Read(buf)
	if err != nil {
		return false
	}

	contentType := http.DetectContentType(buf)

	return contentType == "image/png" || contentType == "image/jpeg"
}
