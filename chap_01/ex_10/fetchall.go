package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/enescakir/emoji"
)

func main() {
	start := time.Now()
	ch := make(chan string) // create a channel - communication mechanism for goroutine
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	//open file to add content to it
	f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("an error occurred: %v\n", err)
		os.Exit(1) // if we can't open file, there's nothing else we can do so let's exit
	}
	defer f.Close() // close file after all content is added
	for range os.Args[1:] {
		msg := <-ch
		if _, err = f.WriteString(msg); err != nil {
			fmt.Printf("an error occurred: %v\n", err)
		}
	}

	fmt.Printf("\nDone %v\n%.2fs elapsed\n", emoji.OkHand.Tone(emoji.Light), time.Since(start).Seconds())

}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("an error occurred: %v\n", err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%7d\t\t%s\n", secs, nbytes, url)
}
