// display command line arguments in key-value form
package main

import (
	"fmt"
	"os"
)

// starts program execution
func main() {
	for k, v := range os.Args {
		fmt.Println(k, "-", v)
	}
}
