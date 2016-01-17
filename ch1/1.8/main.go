// Exercise 1.8: Modify fetch to add the prefix http:// to each argument URL if it is missing. You may want to use strings.HasPrefix.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	const (
		url_prefix = "http://"
	)
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, url_prefix) {
			url = url_prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
