package main

import (
	"io/ioutil"
	"testing"
)

func BenchmarkOcrBytes(b *testing.B) {
	// read result png into buffer
	picBytes, err := ioutil.ReadFile("result.png")
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		_, err := ocrBytes(picBytes)
		if err != nil {
			b.Fatal(err)
		}
	}
}
