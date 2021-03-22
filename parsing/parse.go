package htmlwayfind

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/net/html"
)

func ParseHTML(file io.Reader) *html.Node {
	return nil
}

func ReadFile(folder string) io.Reader {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current working directory", err)
	}
	filePath := filepath.Join(pwd, folder)

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading in file", err)
	}

	return bytes.NewReader(b)
}
