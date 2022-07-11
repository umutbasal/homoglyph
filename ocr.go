package main

import (
	"github.com/otiai10/gosseract/v2"
)

func ocr(path string) string {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(path)
	client.SetLanguage("eng")
	text, err := client.Text()
	if err != nil {
		panic(err)
	}
	return text
}
