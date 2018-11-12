// display all command line argument(s)

package main

import (
	"fmt"
	"os"
	"strings"
)

// starts program execution
func main() {
    fmt.Println(strings.Join(os.Args[1:], " "))
}
