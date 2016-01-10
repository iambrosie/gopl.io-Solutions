// Exercise 1.3: Experiment to measure the difference in running time between our potential inefficient versions and the one that uses Strings.Join.

// Run with:
//   $go test -bench=.
package echo

import (
	"strings"
)

// https://github.com/adonovan/gopl.io/tree/master/ch1/echo1
func FirstEcho(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
}

// https://github.com/adonovan/gopl.io/tree/master/ch1/echo2
func SecondEcho(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
}

// https://github.com/adonovan/gopl.io/tree/master/ch1/echo3
func ThirdEcho(args []string) {
	strings.Join(args[1:], " ")
}
