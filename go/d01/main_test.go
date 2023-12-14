package main

import "testing"

func BenchmarkMain(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mainArgs(false)
	}
}

func BenchmarkRun(b *testing.B) {
	lines := setup()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		run(lines)
	}
}

func BenchmarkP1(b *testing.B) {
	lines := setup()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		runP1(lines)
	}
}

func BenchmarkP2(b *testing.B) {
	lines := setup()
	check := 54249
	var res int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		res = runP2(lines)
	}
	if res != check {
		b.Fail()
	}

}
