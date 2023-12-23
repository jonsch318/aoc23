package main

import (
	"log"
	"testing"
)

const TEST_PATH = "/home/jonas/src/aoc23/input/d06/test"
const INPUT_PATH = "/home/jonas/src/aoc23/input/d06/input"

func TestDefault(t *testing.T) {
	switchImpl(Default)
	check := int64(3316275)
	td, length := setup32(INPUT_PATH)
	log.Printf("--- DEFAULT ---")
	log.Printf("Input %v %v", length, td)
	res := solve32(td, length)
	if res != check {
		log.Printf("%v != %v", res, check)
		t.Fail()
	}

}

func TestAVX(t *testing.T) {
	switchImpl(AVX)
	check := int64(3316275)
	td, length := setup32(INPUT_PATH)
	log.Printf("--- AVX ---")
	log.Printf("Input %v %v", length, td)
	res := solve32(td, length)
	if res != check {
		log.Printf("%v != %v", res, check)
		t.Fail()
	}
}

func TestAVX2(t *testing.T) {
	switchImpl(AVX2)
	check := int64(3316275)
	td, length := setup64(INPUT_PATH)
	log.Printf("--- AVX ---")
	log.Printf("Input %v %v", length, td)
	res := solve64(td, length)
	if res != check {
		log.Printf("%v != %v", res, check)
		t.Fail()
	}
}

func BenchmarkSolveP1AVX(b *testing.B) {
	switchImpl(AVX)
	check := int64(3316275)
	td, length := setup32(INPUT_PATH)

	b.ResetTimer()
	res := int64(0)
	for i := 0; i < b.N; i++ {
		res = solve32(td, length)
	}
	if res != check {
		log.Printf("%v != %v", res, check)
		b.Fail()
	}

}

func BenchmarkSolveP1AVXRead(b *testing.B) {
	switchImpl(AVX)
	check := int64(3316275)

	res := int64(0)
	for i := 0; i < b.N; i++ {
		td, length := setup32(INPUT_PATH)
		res = solve32(td, length)
	}
	if res != check {
		log.Printf("%v != %v", res, check)
		b.Fail()
	}

}

func BenchmarkSolveP1AVX2(b *testing.B) {
	switchImpl(AVX2)
	check := int64(3316275)
	td, length := setup64(INPUT_PATH)

	b.ResetTimer()
	res := int64(0)
	for i := 0; i < b.N; i++ {
		res = solve64(td, length)
	}
	if res != check {
		log.Printf("%v != %v", res, check)
		b.Fail()
	}
}

func BenchmarkSolveP1SCALAR(b *testing.B) {
	switchImpl(Default)
	check := int64(3316275)
	td, length := setup32(INPUT_PATH)

	b.ResetTimer()
	res := int64(0)
	for i := 0; i < b.N; i++ {
		res = solve32(td, length)
	}
	if res != check {
		log.Printf("%v != %v", res, check)
		b.Fail()
	}

}

func BenchmarkP2(b *testing.B) {
	switchImpl(Default)
	check := int64(27102791)
	td, length := setup64("/home/jonas/src/aoc23/input/d06/input")

	b.ResetTimer()
	res := int64(0)
	for i := 0; i < b.N; i++ {
		res = solveP2(td, length)
	}
	if res != check {
		log.Printf("%v != %v", res, check)
		b.Fail()
	}

}
