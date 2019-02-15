
package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	fmt.Println("Unixtime: ", t.Unix())

	// Reference date is "Mon Jan 2 15:04:05 -0700 MST 2006"
	fmt.Println("Format(): ", t.Format("02/01/06-15:04:05"))
}

// Output:
// Unixtime:  1550223175
// Format():  15/02/19-10:32:55
