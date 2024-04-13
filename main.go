package main

import (
	"fmt"
	"github.com/husainbangsat/effectivego/url"
)

func main() {
	fmt.Println("URL Package")
	u := &url.URL{
		Scheme: "https",
		Host:   "foo.com",
		Path:   "go",
	}
	fmt.Println(u)
}
