
package main

import (
	"fmt"
	"flag"
)

func main() {

	// i will be a pointer to int
	var i = flag.Int("i", 1, "Some integer value")

	var s string
	flag.StringVar(&s, "s", "", "Some string value")

	flag.Parse()
	rest := flag.Args()

	fmt.Printf("i: %d\n", *i)
	fmt.Printf("s: %s\n", s)
	for n, arg := range rest {
		fmt.Printf("%d: %s\n", n, arg)
	}
}
