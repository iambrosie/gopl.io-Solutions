// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Key struct {
	line, file string
}

func main() {
	counts := make(map[Key]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for key, n := range counts {
		if n > 1 {
			fmt.Printf("%q \t %q \t %d\n", key.file, key.line, n)
		}
	}
}

func countLines(f *os.File, counts map[Key]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[Key{input.Text(), f.Name()}]++
	}
}
