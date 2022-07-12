package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/chromedp/chromedp"
)

func renderWorker() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, strings.TrimSpace(`
		<html>
		<body>
			<span style="font-size:42px">111</span>
		</body>
		</html>
			`))
	}))

	allocator, cancelEx := chromedp.NewExecAllocator(context.Background(), append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", true))...)
	defer cancelEx()

	ctx, cancel := chromedp.NewContext(allocator)
	defer cancel()
	chromedp.Run(ctx, chromedp.Navigate(ts.URL))

	for text := range renderChan {
		// capture screenshot of an element
		var buf []byte
		if err := chromedp.Run(ctx, putTextAndScreenshot(ts.URL, `span`, text, &buf)); err != nil {
			log.Fatal(err)
		}

		ocrChan <- &buf
	}

}

func putTextAndScreenshot(urlstr, sel, text string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.EvaluateAsDevTools(`document.querySelector("`+sel+`").innerText =`+`'`+text+`'`, nil),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}
