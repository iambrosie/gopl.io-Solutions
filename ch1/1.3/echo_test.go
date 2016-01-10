package echo

import (
	"strings"
	"testing"
)

func build_args(count int) []string {
	var ret []string
	ret = strings.Fields(strings.Repeat("arg ", count))
	return ret
}

func benchmarkFirstEcho(i int, b *testing.B) {
	args := build_args(i)
	for n := 0; n < b.N; n++ {
		FirstEcho(args)
	}
}

func benchmarkSecondEcho(i int, b *testing.B) {
	args := build_args(i)
	for n := 0; n < b.N; n++ {
		SecondEcho(args)
	}
}

func benchmarkThirdEcho(i int, b *testing.B) {
	args := build_args(i)
	for n := 0; n < b.N; n++ {
		ThirdEcho(args)
	}
}

func BenchmarkFirstEcho100(b *testing.B)  { benchmarkFirstEcho(100, b) }
func BenchmarkSecondEcho100(b *testing.B) { benchmarkSecondEcho(100, b) }
func BenchmarkThirdEcho100(b *testing.B)  { benchmarkThirdEcho(100, b) }

func BenchmarkFirstEcho1000(b *testing.B)  { benchmarkFirstEcho(1000, b) }
func BenchmarkSecondEcho1000(b *testing.B) { benchmarkSecondEcho(1000, b) }
func BenchmarkThirdEcho1000(b *testing.B)  { benchmarkThirdEcho(1000, b) }

func BenchmarkFirstEcho10000(b *testing.B)  { benchmarkFirstEcho(10000, b) }
func BenchmarkSecondEcho10000(b *testing.B) { benchmarkSecondEcho(10000, b) }
func BenchmarkThirdEcho10000(b *testing.B)  { benchmarkThirdEcho(10000, b) }
