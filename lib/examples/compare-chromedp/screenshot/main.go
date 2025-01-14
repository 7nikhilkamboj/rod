package main

import (
	"io/ioutil"

	"github.com/7nikhilkamboj/rod"
	"github.com/7nikhilkamboj/rod/lib/proto"
)

// This example demonstrates how to take a screenshot of a specific element and
// of the entire browser viewport, as well as using `kit`
// to store it into a file.
func main() {
	browser := rod.New().MustConnect()

	// capture screenshot of an element
	browser.MustPage("https://google.com").MustElement("#main").MustScreenshot("elementScreenshot.png")

	// capture entire browser viewport, returning jpg with quality=90
	buf, err := browser.MustPage("https://brank.as/").Screenshot(true, &proto.PageCaptureScreenshot{
		Format:  proto.PageCaptureScreenshotFormatJpeg,
		Quality: 90,
	})
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("fullScreenshot.png", buf, 0644)
	if err != nil {
		panic(err)
	}
}
