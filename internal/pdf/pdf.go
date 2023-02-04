package pdf

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

// Converts input image path to pdf in output directory
func ConvertImage(inputFilePath, outputDir string) error {
	// images/test.jpg => test.jpg
	imageFileName := path.Base(inputFilePath)
	// test.jpg => test.pdf
	newFilename := strings.Replace(imageFileName, filepath.Ext(imageFileName), ".pdf", -1)
	// outputdir/test.pdf
	outputFilePath := path.Join(outputDir, newFilename)

	pdf := gofpdf.New(gofpdf.OrientationPortrait, "mm", gofpdf.PageSizeLetter, "")
	pdf.AddPage()
	pdf.Image(inputFilePath, 0, 0, 240, 480, false, "", 0, "")

	err := pdf.OutputFileAndClose(outputFilePath)
	if err != nil {
		return err
	}

	return nil
}

// Finds all of the generated pdfs in the output dir
func Find(dir string) ([]string, error) {
	filePaths := []string{}

	files, err := os.ReadDir(dir)
	if err != nil {
		return filePaths, err
	}

	for _, file := range files {
		filePath := path.Join(dir, file.Name())
		filePaths = append(filePaths, filePath)
	}

	return filePaths, nil
}
