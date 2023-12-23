package main

import (
	"log"
	"testing"
)

const TEST_PATH = "/home/jonas/src/aoc23/input/d06/test"
const INPUT_PATH = "/home/jonas/src/aoc23/input/d06/input"

func TestMain(t *testing.T) {
	switchImpl(Default)

	check := int64(3316275)
	td, length := setup(INPUT_PATH)
	log.Printf("--- DEFAULT ---")
	log.Printf("Input %v %v", length, td)
	res := solveP1(td, length)
	if res != check {
		log.Printf("%v != %v", res, check)
		t.Fail()
	}

	switchImpl(AVX)

	check = int64(3316275)
	td, length = setup(INPUT_PATH)
	log.Printf("--- AVX ---")
	log.Printf("Input %v %v", length, td)
	res = solveP1(td, length)
	if res != check {
		log.Printf("%v != %v", res, check)
		t.Fail()
	}
}

func BenchmarkSolveP1AVX(b *testing.B) {
	switchImpl(AVX)
	check := int64(3316275)
	td, length := setup(INPUT_PATH)

	b.ResetTimer()
	res := int64(0)
	for i := 0; i < b.N; i++ {
		res = solveP1(td, length)
	}
	if res != check {
		log.Printf("%v != %v", res, check)
		b.Fail()
	}

}

func BenchmarkSolveP1SCALAR(b *testing.B) {
	switchImpl(Default)
	check := int64(3316275)
	td, length := setup(INPUT_PATH)

	b.ResetTimer()
	res := int64(0)
	for i := 0; i < b.N; i++ {
		res = solveP1(td, length)
	}
	if res != check {
		log.Printf("%v != %v", res, check)
		b.Fail()
	}

}
