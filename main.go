package main

import "fmt"

func main() {
	render("ЕＬоｎ ᛖ⋃ѕkｏ", "result.png")
	fmt.Println(ocr("result.png"))
}
