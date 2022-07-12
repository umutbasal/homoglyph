package main

import (
	"fmt"
	"log"

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

func ocrWorker() {
	for picBytes := range ocrChan {
		log.Println("ocrchan got picBytes")
		text, err := ocrBytes(*picBytes)
		if err != nil {
			fmt.Println(err)
		}
		log.Printf("ocr resolved for byte len: %d, text: %s", len(*picBytes), text)
	}
}
