package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"unsafe"

	solve "github.com/jonsch318/aoc23/go/d06/solve"
)

var impl = Default

const (
	Default int = iota
	AVX
	AVX2
)

func switchImpl(nImpl int) {
	impl = nImpl
}

func main() {
	races, length := setup("/home/jonas/src/aoc23/input/d06/input")
	log.Printf("SETUP: %v", races)
	res := solveP1(races, length)
	log.Printf("P1: %v", res)
}

func solveP1(timesDistances []float32, length int) int64 {
	prod := int64(1)
	switch impl {
	case AVX:
		for i := 0; i < len(timesDistances); i += 8 {
			ptr := solve.SolveAVX([4]float32(timesDistances[i:i+4]), [4]float32(timesDistances[i+4:i+8]))
			prod *= int64(ptr[0])
			prod *= int64(ptr[1])
			prod *= int64(ptr[2])
			prod *= int64(ptr[3])
		}
	case AVX2:
		for i := 0; i < len(timesDistances); i += 16 {
			prod *= int64(solve.SolveAVX2(unsafe.Pointer(&timesDistances[i]), unsafe.Pointer(&timesDistances[i+8]))) // Use the Solve function from the imported package
		}
	case Default:
		for i := 0; i < len(timesDistances); i += 2 {
			prod *= int64(solve.SolveScalar(timesDistances[i], timesDistances[i+1]))
		}
	}
	return prod
}

func toDouble(a string) float32 {
	res, _ := strconv.ParseFloat(a, 32)
	return float32(res)
}

func setup(path string) ([]float32, int) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Panicf("not good %v", err)
		return nil, 0
	}
	text := string(data)

	lines := strings.Split(text, "\n")
	lines = lines[0:2]

	timeStrings := strings.Fields(lines[0])[1:]
	distanceStrings := strings.Fields(lines[1])[1:]

	// cache efficient array. If AVX2 => [8 times...][8 distances...][8 times...]... else
	// Default => [time][distance][time]...

	var vectorLength = 1
	length := len(timeStrings)
	switch impl {
	case AVX:
		vectorLength = 128 / (4 * 8)
	case AVX2:
		vectorLength = 256 / (4 * 8)
	}

	if length%vectorLength != 0 {
		length += vectorLength - (len(timeStrings) % vectorLength)
	}

	timesDistances := make([]float32, length*2)
	for i := 0; i < len(timeStrings); i += vectorLength {
		for j := 0; j < vectorLength; j++ {
			x := (i * 2) + j
			if len(timeStrings) <= i+j {
				timesDistances[x] = 0
				timesDistances[x+vectorLength] = 0
			} else {
				timesDistances[x] = toDouble(timeStrings[i+j])
				timesDistances[x+vectorLength] = toDouble(distanceStrings[i+j])
			}
		}
	}
	return timesDistances, len(timeStrings)
}
