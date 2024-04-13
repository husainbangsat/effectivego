package url_test

import (
	"fmt"
	"github.com/husainbangsat/effectivego/url"
	"log"
)

func ExampleURL() {
	u, err := url.Parse("https://foo.com/go")
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	fmt.Println(u)
	// Output:
	// https://foo.com/go
}
