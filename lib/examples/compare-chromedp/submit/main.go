package main

import (
	"log"
	"strings"

	"github.com/7nikhilkamboj/rod"
	"github.com/7nikhilkamboj/rod/lib/input"
)

//This example demonstrates how to fill out and submit a form.
func main() {
	page := rod.New().MustConnect().MustPage("https://github.com/search")

	page.MustElement(`input[name=q]`).MustWaitVisible().MustInput("chromedp").MustPress(input.Enter)

	res := page.MustElementR("a", "chromedp").MustParent().MustNext().MustText()

	log.Printf("got: `%s`", strings.TrimSpace(res))
}
