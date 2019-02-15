
package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {

	u := "https://foo.de/bar/1?q=somevalue&search=none&q=bar"
	url, err := url.Parse(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Opaque:", url.Opaque)
	fmt.Println("Host:", url.Host)
	fmt.Println("Path:", url.Path)
	fmt.Println("RawPath:", url.RawPath)
	fmt.Println("RawQuery:", url.RawQuery)
	q := url.Query()
	for key, list := range q {
		fmt.Printf("    %s: [%s]\n", key, strings.Join(list, ", "))
	}
}

