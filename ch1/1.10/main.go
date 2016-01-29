// Exercise 1.10: Find a website that produces a large amount of data. Investigate caching by running fetchall twice in succession to see whether the reported time changes much. Do you get the same content each time? Modify fetchall to print its output to a file so that it can be examined.

// output is saved in the local folder, in a file called findall.out
// i ran it with the following command:
//		go run main.go http://www.lagado.com/tools/cache-test https://www.drupal.org https://www.google.com
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	out, err := os.OpenFile("findall.out", os.O_RDWR|os.O_APPEND|os.O_CREATE, os.FileMode(0666))
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	defer out.Close()
	for range os.Args[1:] {
		_, err := out.WriteString(<-ch + "\n")
		if err != nil {
			ch <- fmt.Sprint(err) // send to channel ch
			return
		}
	}
	out.WriteString(time.Since(start).String() + " elapsed\n\n")
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
