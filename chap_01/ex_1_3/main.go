// benchmarking the different ways of diplaying/retrieving command line argument
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// times a process
func timeTrack(start time.Time, name string) {
	elpased := time.Since(start)
	log.Printf("%s took: %s", name, elpased)
	fmt.Println()

}

// starts program execution
func main() {
	start := time.Now()
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
	timeTrack(start, "'loop'")

	start = time.Now()
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
	timeTrack(start, "'for+range'")

	start = time.Now()
	fmt.Println(strings.Join(os.Args, "\n"))
	timeTrack(start, "'join'")

}
