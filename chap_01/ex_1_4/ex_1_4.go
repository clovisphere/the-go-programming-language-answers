package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
        fmt.Println("Type <q> to end program")
		countLines(os.Stdin)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f)
			f.Close()
		}
	}
}


func countLines(f *os.File) {
    counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
        // gracefully break out of loop
        if input.Text() == "q" {
            break
        }
		counts[input.Text()]++
	}
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%s\t%d\t%s\n", f.Name(), n, line)
        }
    }
}



