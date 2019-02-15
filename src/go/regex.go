
package main

import (
	"fmt"
	"regexp"
)

func main() {

	s := "asdf{} asdf{} foo{}"

	re := regexp.MustCompile("{}")
	l := re.FindAllString(s, -1)

	for _, m := range l {
		fmt.Println(m)
	}
}

// Output:
// {}
// {}
// {}
