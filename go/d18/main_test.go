package main

import (
	"testing"

	"github.com/jonsch318/aoc23/go/d18/solve"
)

const ABS_INPUT = "/home/jonas/src/aoc23/input/d18/input"

func BenchmarkP1(b *testing.B) {
	lines := setup(ABS_INPUT, readP1)
	check := int64(92758)
	b.ResetTimer()
	var res int64
	for n := 0; n < b.N; n++ {
		res = solve.Solve(lines)
	}

	if res != check {
		b.Fail()
	}
}

func BenchmarkP2(b *testing.B) {
	lines := setup(ABS_INPUT, readP2)
	check := int64(62762509300678)
	var res int64
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		res = solve.Solve(lines)
	}
	if res != check {
		b.Fail()
	}

}
