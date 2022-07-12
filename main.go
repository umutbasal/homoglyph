package main

import (
	"math/rand"
)

var renderChan = make(chan string)
var ocrChan = make(chan *[]byte)

func main() {
	go renderWorker()
	go ocrWorker()

	testCharSet := `A Æ B Ɓ Ƃ C Ƈ D Đ Ɖ Ɗ Ƌ Ð E Ǝ Ə Ɛ F Ƒ G Ǥ Ɠ Ɣ Ƣ H Ƕ Ħ I Ɨ Ɩ J K Ƙ L Ł M N Ɲ Ƞ Ŋ O Œ Ø Ɔ Ɵ Ȣ P Ƥ Q R Ʀ S ß Ʃ T Ŧ Ƭ Ʈ U Ɯ Ʊ V Ʋ W X Y Ƴ Z Ƶ Ȥ Ʒ Ƹ Ȝ Þ Ƿ Ƨ Ƽ Ƅ`
	randomWords := []string{"ЕＬоｎ ᛖ⋃ѕkｏ"}

	// randomly create 3-4 words from charset
	for i := 0; i < 30; i++ {
		randomWords = append(randomWords, randomWord(testCharSet))
	}

	for _, text := range randomWords {
		renderChan <- text
	}
}

func randomWord(charset string) string {
	word := ""
	for i := 0; i < 6; i++ {
		char := string(charset[rand.Intn(len(charset))])
		if char == "" || char == " " {
			continue
		}
		word += char
	}
	return word
}
