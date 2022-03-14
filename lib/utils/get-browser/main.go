package main

import (
	"fmt"

	"github.com/7nikhilkamboj/rod/lib/launcher"
	"github.com/7nikhilkamboj/rod/lib/utils"
)

func main() {
	p, err := launcher.NewBrowser().Get()
	utils.E(err)

	fmt.Println(p)
}
