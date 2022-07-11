package main

import (
	"github.com/otiai10/gosseract/v2"
)

func ocr(path string) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	client.SetImage(path)
	client.SetLanguage("eng")
	return client.Text()
}

func ocrBytes(picBytes []byte) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	client.SetImageFromBytes(picBytes)
	client.SetLanguage("eng")
	return client.Text()
}
