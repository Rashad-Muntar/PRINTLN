package utils

import (
	"io/ioutil"
	"bytes"
	"os"
	"log"
	"github.com/unidoc/unipdf/v3/model"
)

func ProcessPDF(filename string) (int, error) {
		// Open the PDF file
		file, err := os.Open(filename)
		if err != nil {
			log.Printf("Error opening PDF file: %s\n", err.Error())
			return 0, err
		}
		defer file.Close()
	
		// Read the file contents into memory
		buf, err := ioutil.ReadAll(file)
		if err != nil {
			log.Printf("Error reading PDF file: %s\n", err.Error())
			return 0, err
		}
	
		// Create an io.ReadSeeker from the buffer
		reader := bytes.NewReader(buf)
	
		// Create a PDF reader
		pdf, err := model.NewPdfReader(reader)
		if err != nil {
			log.Printf("Error creating PDF reader: %s\n", err.Error())
			return 0, err
		}
	
		// Get the total number of pages
		pagesCount, err := pdf.GetNumPages()
		if err != nil {
			log.Printf("Error getting number of PDF pages: %s\n", err.Error())
			return 0, err
		}
	
		return pagesCount, nil
}
